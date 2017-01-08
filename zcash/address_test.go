package zcash

import "testing"

var golden = []struct{ skey, addr string }{
	{"SKxt8pwrQipUL5KgZUcBAqyLj9R1YwMuRRR3ijGMCwCCqchmi8ut", "zcJLC7a3aRJohMNCVjSZQ8jFuofhAHJNAY4aX5soDkYfgNejzKnEZbucJmVibLWCwK8dyyfDhNhf3foXDDTouweC382LcX5"},
	{"SKxoo5QkFQgTbdc6EWRKyHPMdmtNDJhqudrAVhen9b4kjCwN6CeV", "zcRYvLiURno1LhXq95e8avXFcH2fKKToSFfhqaVKTy8mGH7i6SJbfuWcm4h9rEA6DvswrbxDhFGDQgpdDYV8zwUoHvwNvFX"},
	{"SKxsVGKsCESoVb3Gfm762psjRtGHmjmv7HVjHckud5MnESfktUuG", "zcWGguu2UPfNhh1ygWW9Joo3osvncsuehtz5ewvXd78vFDdnDCRNG6QeKSZpwZmYmkfEutPVf8HzCfBytqXWsEcF2iBAM1e"},
	{"SKxp72QGQ2qtovHSoVnPp8jRFQpHBhG1xF8s27iRFjPXXkYMQUA6", "zcWZomPYMEjJ49S4UHcvTnhjYqogfdYJuEDMURDpbkrz94bkzdTdJEZKWkkpQ8nK62eyLkZCvLZDFtLC2Cq5BmEK3WCKGMN"},
	{"SKxpmLdykLu3xxSXtw1EA7iLJnXu8hFh8hhmW1B2J2194ijh5CR4", "zcgjj3fJF59QGBufopx3F51jCjUpXbgEzec7YQT6jRt4Ebu5EV3AW4jHPN6ZdXhmygBvQDRJrXoZLa3Lkh5GqnsFUzt7Qok"},
}

func TestKeyToAddress(t *testing.T) {
	for i, g := range golden {
		key, version, err := Base58Decode(g.skey)
		if err != nil {
			t.Fatal(i, err)
		}
		if version != ProdSpendingKey {
			t.Fatal(i, version)
		}
		rawAddr, err := KeyToAddress(key)
		if err != nil {
			t.Fatal(i, err)
		}
		ourAddr := Base58Encode(rawAddr, ProdAddress)
		if g.addr != ourAddr {
			t.Errorf("%d: addr %s, want %s", ourAddr, g.addr)
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
