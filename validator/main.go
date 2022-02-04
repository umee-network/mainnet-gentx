package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	umeeapp "github.com/umee-network/umee/app"
)

func main() {
	for i, file := range os.Args {
		if i == 0 {
			continue
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
		genStateBytes := encCfg.Marshaler.MustMarshalJSON(&genState)

		err = umeeapp.GenutilModule{}.ValidateGenesis(encCfg.Marshaler, encCfg.TxConfig, genStateBytes)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("gentx is valid")
	}

}
