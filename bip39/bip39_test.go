package bip39

import (
	"bytes"
	crnd "crypto/rand"
	"encoding/hex"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestDivideUp(t *testing.T) {
	if divideUp(10, 2) != 5 {
		t.Fail()
	}
	if divideUp(9, 2) != 5 {
		t.Fail()
	}
	if divideUp(10, 5) != 2 {
		t.Fail()
	}
	if divideUp(8, 5) != 2 {
		t.Fail()
	}
}

func TestRandomBip39(t *testing.T) {
	for i := 1; i < 1000; i++ {
		num := rand.Intn(100) * 32 / 8
		data := make([]byte, num)
		if _, err := crnd.Read(data); err != nil {
			t.Fatal(err)
		}
		mnemonic := Encode(data)
		res, c, err := Decode(mnemonic)
		if err != nil {
			t.Log("\n" + hex.Dump(data))
			t.Fatal(err)
		}
		if c != nil {
			t.Fatal(c)
		}
		if !bytes.Equal(res, data) {
			t.Log("\n" + hex.Dump(data))
			t.Log("\n" + hex.Dump(res))
			t.Fail()
		}
	}
}

func TestBip39Mistakes(t *testing.T) {
	goodData, c, err := Decode([]string{"enter", "evidence", "garage"})
	if err != nil {
		t.Fatal(err)
	}
	if c != nil {
		t.Fatal(c)
	}

	data, c, err := Decode([]string{"enter", "evidente", "garage"})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, goodData) {
		t.Log("\n" + hex.Dump(data))
		t.Log("\n" + hex.Dump(goodData))
		t.Fail()
	}
	if !reflect.DeepEqual(c, []string{"evidente -> evidence"}) {
		t.Fatal(c)
	}

	data, c, err = Decode([]string{"enter", "evidente", "garale"})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, goodData) {
		t.Log("\n" + hex.Dump(data))
		t.Log("\n" + hex.Dump(goodData))
		t.Fail()
	}
	if !reflect.DeepEqual(c, []string{
		"evidente -> evidence", "garale -> garage"}) {
		t.Fatal(c)
	}

	_, _, err = Decode([]string{"access", "chronic", "cricket", "search", "magnet", "myself", "peasant", "party", "party"})
	if err == nil {
		t.Fatal("expected checksum error")
	}
}
