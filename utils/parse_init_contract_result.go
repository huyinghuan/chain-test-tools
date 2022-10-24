package utils

import "math/big"

//
func ParseInitConTractResult(body []byte) []interface{} {
	index := uint64(0)
	length := uint64(len(body))
	values := []interface{}{}
	for index < length {

		if body[index] == byte(10) {
			if index+1 < length {
				// 读取内容长度
				index = index + 1
				vLength := new(big.Int).SetBytes(body[index : index+1]).Uint64()
				index = index + 1
				if index+vLength > length {
					break
				}
				v := string(body[index : index+vLength])
				values = append(values, v)
				index = index + vLength
				continue
			}
			break
		}

		if body[index] == byte(18) {
			if index+1 < length {
				// 读取内容长度
				index = index + 1
				vLength := new(big.Int).SetBytes(body[index : index+1]).Uint64()
				index = index + 1
				if index+vLength > length {
					break
				}
				v := string(body[index : index+vLength])
				values = append(values, v)
				index = index + vLength
				continue
			}
			break
		}

		if body[index] == byte(7) {
			if index+6 < length {
				index = index + 1
				if string(body[index:index+5]) == "-----" {
					values = append(values, string(body[index:length-1]))
					break
				}
				continue
			}
			break
		}
		index = index + 1
	}
	return values

}
