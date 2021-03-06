package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	gravitytypes "github.com/Gravity-Bridge/Gravity-Bridge/module/x/gravity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	umeeapp "github.com/umee-network/umee/app"
)

var (
	minCommissionAllowed = sdk.MustNewDecFromStr("0.02")
)

func main() {
	for i, file := range os.Args {
		if i == 0 {
			continue
		}

		if !strings.Contains(file, ".json") {
			continue
		}

		if filepath.Dir(file) != "gentxs" {
			log.Fatal("gentxs should go into the gentx directory")
		}

		filecontents, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		var gentx json.RawMessage
		err = json.Unmarshal(filecontents, &gentx)
		if err != nil {
			log.Fatal(err)
		}

		encCfg := umeeapp.MakeEncodingConfig()
		genState := genutiltypes.GenesisState{GenTxs: []json.RawMessage{gentx}}

		txJSONDecoder := encCfg.TxConfig.TxJSONDecoder()
		for i, genTx := range genState.GenTxs {
			var tx sdk.Tx

			tx, err := txJSONDecoder(genTx)
			if err != nil {
				log.Fatal(err)
			}

			msgs := tx.GetMsgs()
			if n := len(msgs); n != 2 {
				log.Fatal(fmt.Errorf(
					"gentx %d contains invalid number of messages; expected: 2; got: %d",
					i, n,
				))
			}

			if msgCreateVal, ok := msgs[0].(*stakingtypes.MsgCreateValidator); ok {
				err := msgCreateVal.ValidateBasic()
				if err != nil {
					log.Fatal(err)
				}

				if msgCreateVal.Value.Denom != umeeapp.BondDenom {
					log.Fatalf("Delegation denomination must be %s", umeeapp.BondDenom)
				}

				if msgCreateVal.Commission.Rate.LT(minCommissionAllowed) {
					log.Fatalf("Validator commission must be at least 2%%: %s", msgCreateVal.Commission.Rate)
				}
			} else {
				log.Fatal(fmt.Errorf(
					"gentx %d contains invalid message at index 0; expected: %T; got: %T",
					i, &stakingtypes.MsgCreateValidator{}, msgs[0],
				))
			}

			if msgSetOrchAddr, ok := msgs[1].(*gravitytypes.MsgSetOrchestratorAddress); ok {
				err := msgSetOrchAddr.ValidateBasic()
				if err != nil {
					log.Fatal(err)
				}
			} else {
				log.Fatal(fmt.Errorf(
					"gentx %d contains invalid message at index 1; expected: %T; got: %T",
					i, &gravitytypes.MsgSetOrchestratorAddress{}, msgs[1],
				))
			}

		}

		// double check it's a well formed TX
		tx, err := encCfg.TxConfig.TxJSONDecoder()(filecontents)
		if err != nil {
			log.Fatal(err)
		}
		err = tx.ValidateBasic()
		if err != nil {
			log.Fatal(err)
		}

		txBuilder, err := encCfg.TxConfig.WrapTxBuilder(tx)
		if err != nil {
			log.Fatal(err)
		}

		signatures, err := txBuilder.GetTx().GetSignaturesV2()
		if err != nil {
			log.Fatal(err)
		}

		// validate signatures
		for _, sig := range signatures {
			err := signing.VerifySignature(sig.PubKey, signing.SignerData{
				ChainID:       "umee-1",
				AccountNumber: 0,
				Sequence:      sig.Sequence,
			},
				sig.Data,
				encCfg.TxConfig.SignModeHandler(),
				tx,
			)
			if err != nil {
				log.Fatal(err)
			}
		}

		// this is a bit redundant, but it doesn't hurt to run it
		genStateBytes := encCfg.Marshaler.MustMarshalJSON(&genState)
		err = umeeapp.GenutilModule{}.ValidateGenesis(encCfg.Marshaler, encCfg.TxConfig, genStateBytes)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("gentx is valid")
	}
}
