validate-gentxs:
	cd validator; go build -o ../gentxvalidator; cd ..;  ./gentxvalidator $(GIT_DIFF)