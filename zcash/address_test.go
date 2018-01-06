package zcash

import (
	"bytes"
	"testing"
)

var golden = []struct{ skey, addr, vk string }{
	{"SKxt8pwrQipUL5KgZUcBAqyLj9R1YwMuRRR3ijGMCwCCqchmi8ut", "zcJLC7a3aRJohMNCVjSZQ8jFuofhAHJNAY4aX5soDkYfgNejzKnEZbucJmVibLWCwK8dyyfDhNhf3foXDDTouweC382LcX5", "ZiVKYQyUcyAJLKwcosSeDxkGRhygFdAPWsr3m8UgjC5X85yqNyLTtJJJYNH83Wf2AQKU6TZsd65MXBZLFj6eSCAFcnCFuVCFS"},
	{"SKxoo5QkFQgTbdc6EWRKyHPMdmtNDJhqudrAVhen9b4kjCwN6CeV", "zcRYvLiURno1LhXq95e8avXFcH2fKKToSFfhqaVKTy8mGH7i6SJbfuWcm4h9rEA6DvswrbxDhFGDQgpdDYV8zwUoHvwNvFX", "ZiVKfdhhmQ1fpXaxyW5zRXw4Dhg9cbKRgK7mNFoBLiKjiBZiHJYJTpV2gNMDMPY9sRC96vnKZcnTMSi65SKPyL4WNQNm9PT5H"},
	{"SKxsVGKsCESoVb3Gfm762psjRtGHmjmv7HVjHckud5MnESfktUuG", "zcWGguu2UPfNhh1ygWW9Joo3osvncsuehtz5ewvXd78vFDdnDCRNG6QeKSZpwZmYmkfEutPVf8HzCfBytqXWsEcF2iBAM1e", "ZiVKkMUGwx4GgtwxTedRHYewVVskWicz8APQgdcYmvUsiLYgSh3cLAa8TwiR3shyNngGbLiUbYMkZ8F1giXmmcED98rDMwNSG"},
	{"SKxp72QGQ2qtovHSoVnPp8jRFQpHBhG1xF8s27iRFjPXXkYMQUA6", "zcWZomPYMEjJ49S4UHcvTnhjYqogfdYJuEDMURDpbkrz94bkzdTdJEZKWkkpQ8nK62eyLkZCvLZDFtLC2Cq5BmEK3WCKGMN", "ZiVKkeb8STw7kpJQsjRCQKovQBciPcfjkpajuuS25DTXSQSVasnq4BkyaMLBBxAkZ8fv6f18woWgaA8W7kGvYp1C1ESaWGjwV"},
	{"SKxpmLdykLu3xxSXtw1EA7iLJnXu8hFh8hhmW1B2J2194ijh5CR4", "zcgjj3fJF59QGBufopx3F51jCjUpXbgEzec7YQT6jRt4Ebu5EV3AW4jHPN6ZdXhmygBvQDRJrXoZLa3Lkh5GqnsFUzt7Qok", "ZiVKvpWQiDpxAvWTMLkjjSbCiBGc4kXhtkgAJfW1JVbCTUY4YaAVvVZzCz6wspG9qttciRFLEXm3HLQAmssFbUp9uPEkP3uu5"},
}

func TestKeyToAddress(t *testing.T) {
	for i, g := range golden {
		key, version, err := Base58Decode(g.skey)
		if err != nil {
			t.Fatal(i, err)
		}
		if bytes.Compare(version[:], ProdSpendingKey) != 0 {
			t.Fatal(i, version)
		}
		rawAddr, err := KeyToAddress(key)
		if err != nil {
			t.Fatal(i, err)
		}

		viewingkey, err := KeyToViewingKey(key)
		if err != nil {
			t.Fatal(i, err)
		}

		ourAddr := Base58Encode(rawAddr, ProdAddress)

		if g.addr != ourAddr {
			t.Errorf("%d: addr %s, want %s", i, ourAddr, g.addr)
		}

		outVK := Base58Encode(viewingkey, ProdViewingKey)

		if g.vk != outVK {
			t.Errorf("%d: addr %s, want %s", i, outVK, g.vk)
		}
	}
}

func TestGenerateKey(t *testing.T) {
	key := GenerateKey()
	_, err := KeyToAddress(key)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateVanityKey(t *testing.T) {
	key := GenerateVanityKey("zcaa", ProdAddress)
	rawAddr, err := KeyToAddress(key)
	if err != nil {
		t.Fatal(err)
	}
	addr := Base58Encode(rawAddr, ProdAddress)
	if addr[:4] != "zcaa" {
		t.Fatal(addr)
	}
}
