package ethex

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

//wei 转换成 eth
func WeiToEth(val string) *big.Float {
	bigval := new(big.Float)
	bignum, _ := bigval.SetString(val)
	amount := big.NewFloat(math.Pow(10, 18.0))
	bigf := new(big.Float).Quo(bignum, amount) // 除法 10的 18次方

	return bigf
}

/*
1kwei = 10^3wei(10的 3次幂)svg

1mwei = 10^6wei(10的 6次幂)xml

1gwei = 10^9wei(10的 9次幂)get

1szabo = 10^12wei(10的 12次幂)it

1finney = 10^15wei(10的 15次幂)比特币

1ether = 10^18wei(10的 18次幂)以太坊

1kether = 10^21wei(10的 21次幂)webkit

1mether = 10^24wei(10的 24次幂)chart

1gether = 10^27wei(10的 27次幂)di
*/
//EthToWei eth单位安全转wei
//https://stackoverrun.com/cn/q/13021596
func EthToWei(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(10 ^ 18))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}

// 0x 16进制 转换 10进制
func Num_16To10Int64(value interface{}) (int64, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		if len(v) > 2 {
			sub := v[:2]
			if sub == "0x" || sub == "0X" {
				return strconv.ParseInt(v[2:], 16, 64)
				//return strconv.ParseInt(v, 0, 0)  这行也可以。
			}
		}

		return strconv.ParseInt(v, 16, 64)
	default:
		return int64(0), errors.New("unknown type")
	}
}

//  10 -> 16 进制 字符串 加 0x
func Num_10To16String(value interface{}) (string, error) {
	switch value.(type) {
	case string:
		v, _ := value.(string)
		num, _ := strconv.ParseInt(v, 10, 64)
		return fmt.Sprintf("0x%x", num), nil
	case int:
		v, _ := value.(int)
		return fmt.Sprintf("0x%x", v), nil
	case int32:
		v, _ := value.(int32)
		return fmt.Sprintf("0x%x", v), nil
	case int64:
		v, _ := value.(int64)
		return fmt.Sprintf("0x%x", v), nil
	case float32:
		v, _ := value.(float32)
		return fmt.Sprintf("0x%x", v), nil
	case float64:
		v, _ := value.(float64)
		return fmt.Sprintf("0x%x", v), nil
	default:
		return "", errors.New("unknown type")
	}
}
