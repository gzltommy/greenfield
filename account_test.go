package greenfield

import "testing"

func TestAccount(t *testing.T) {
	balance, err := GetAccountBalance("0x4bDA26282Cd8D7E5B5253e339d9E7906B039b2e6")
	t.Log(err)
	t.Log(balance)
}
