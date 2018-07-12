package main

import "strconv"

/* Increasing this would frustate CPU */
const targetBits = 18
const subsidy = 25
const blocksBucket = "blocks"
const dbFile = "blocks.db"
const EOFerr = "EOF"
/* Legendary: ** */
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink" +
                                "of second bailout for banks"

func IntToHex(n int64) []byte {
    return []byte(strconv.FormatInt(n, 16))
}
