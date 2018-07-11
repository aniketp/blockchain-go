package main

import "strconv"

/* Increasing this would frustate CPU */
var targetBits int64 = 18

func IntToHex(n int64) []byte {
    return []byte(strconv.FormatInt(n, 16))
}
