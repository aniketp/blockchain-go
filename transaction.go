package main

type Transaction struct {
	ID		[]byte
	Vin		[]TXInput
	Vout		[]TXOutput
}

type TXOutput struct {
	Value		int64
	ScriptPubKey	string
}

type TXInput struct {
	Txid		[]byte
	Vout		int64
	ScriptSig	string
}