package fuzz

import "strconv"
import "github.com/go-numb/go-ftx/auth"
import "github.com/go-numb/go-ftx/types"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) < 1 {
        num = 0
    } else {
        num, _ = strconv.Atoi(string(bytes[0]))
    }
    switch num {
        case 1:
            content := string(bytes)
            var test auth.Config
            test.Signature(content)
            return 0

        default:
            var test types.FtxTime
            test.UnmarshalJSON(bytes)
            return 0
    }
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}