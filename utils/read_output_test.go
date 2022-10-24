package utils

import (
	"chain-api-imgo/resource"
	"testing"
)

func TestReadInputResult(t *testing.T) {
	abibytes, _ := resource.ReadFile("asserts/test/mgtv.abi")
	method, args, _ := ReadInput(string(abibytes), "23b872dd0000000000000000000000000fbabff26673a2bb136e3b4fc589e53918aafa77000000000000000000000000373cc05284857989c4eb8e5adcbbe23bb089b66c000000000000000000000000000000000000000000000000098d6b53e244d000")

	t.Log(method, args)
}
