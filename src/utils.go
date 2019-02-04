package main

import "strconv"

/* Increasing this would frustate CPU */
const targetBits = 18
const blocksBucket = "blocks"
const dbFile = "blocks.db"
const EOFerr = "EOF"

func IntToHex(n int64) []byte {
    return []byte(strconv.FormatInt(n, 16))
}
