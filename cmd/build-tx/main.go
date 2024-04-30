package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	cardano "github.com/echovl/cardano-go"
	"github.com/fxamacker/cbor/v2"
	"gopkg.in/yaml.v2"
)

type TransactionData struct {
	UTXOs            []UTXO `yaml:"utxos"`
	ReceivingAddress string `yaml:"receiving_address"`
	PolicyHash       string `yaml:"policy_hash"`
	AssetName        string `yaml:"asset_name"`
}

type UTXO struct {
	TXID   string `yaml:"txid"`
	Index  int    `yaml:"index"`
	Amount int64  `yaml:"amount"`
}

func main() {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	var config TransactionData
	if err = yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("error unmarshaling config data: %v", err)
	}

	receivingAddress, err := cardano.NewAddress(config.ReceivingAddress)
	if err != nil {
		log.Fatal(err)
	}

	// NOTE: this is not needed if the tx drains the address
	changeAddr := receivingAddress

	policyHash, err := hex.DecodeString(config.PolicyHash)
	if err != nil {
		panic(err)
	}
	policyID := cardano.NewPolicyIDFromHash(policyHash)
	assetName := cardano.NewAssetName(config.AssetName)

	txBuilder := cardano.NewTxBuilder(&cardano.ProtocolParams{})
	txBuilder.AddChangeIfNeeded(changeAddr)

	for _, utxo := range config.UTXOs {
		value := cardano.NewValueWithAssets(
			10e6, // <--- this may be problematic, is this the n decimals of the token in question?
			cardano.NewMultiAsset().Set(policyID, cardano.NewAssets().Set(assetName, 10e6)),
		)

		txBuilder.AddInputs(
			cardano.NewTxInput(cardano.Hash32(utxo.TXID), uint(utxo.Index), value),
		)
	}
	txBuilder.AddOutputs(
		cardano.NewTxOutput(receivingAddress, cardano.NewValueWithAssets(42e6, cardano.NewMultiAsset())),
	)
	tx, err := txBuilder.Build()
	if err != nil {
		err := fmt.Errorf("failure to build tx: %w", err)
		log.Fatal(err)
	}

	bytes, err := tx.MarshalCBOR()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#x\n", bytes)
	var transaction cardano.Tx
	if err := cbor.Unmarshal(bytes, &transaction); err != nil {
		log.Fatal(err)
	}

	for _, input := range transaction.Body.Inputs {
		fmt.Printf("%s:%d\n", string(input.TxHash), input.Index)
	}
	fmt.Printf("%#v\n", transaction.Body.Outputs[0].Amount.MultiAsset)
	for _, k := range transaction.Body.Outputs[0].Amount.MultiAsset.Keys() {
		fmt.Printf("k: %#v; v: %#v\n", k, transaction.Body.Outputs[0].Amount.MultiAsset.Get(k))
	}
	fmt.Printf("%d\n", transaction.Body.Outputs[0].Amount.Coin)
}
