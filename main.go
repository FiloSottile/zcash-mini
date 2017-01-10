package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/FiloSottile/zcash-mini/bip39"
	"github.com/FiloSottile/zcash-mini/zcash"
)

var logo = `

             :88SX@888@@X8:
          %%Xt%tt%SSSSS:XXXt@@
        @S;;tt%%%t    ;;::XXXXSX
      .t:::;;%8888    88888tXXXX8;
     .%...:::8             8::XXX%;
     8888...:t888888X     8t;;::XX8
    %888888...:::;:8    :Xttt;;;::X@
    888888888...:St    8:%%tttt;;;:X
    88888888888S8    :%;ttt%%tttt;;X
    %888888888%t    8S:;;;tt%%%ttt;8
     8t8888888     S8888888Stt%%%t@
     .@tt888@              8;;ttt@;
      .8ttt8@SSSSS    SXXXX%:;;;X;
        X8ttt8888%    %88...::X8
          %8@tt88;8888%8888%8X
             :@888@XXX@888:

                    _       _
          _ __ ___ (_)_ __ (_)
         | '_ \ _ \| | '_ \| |
         | | | | | | | | | | |
         |_| |_| |_|_|_| |_|_|
`

var template = `%s

###############################################################
#
#  Here is your z-address
#
#      %s
#
###############################################################
#
#  Here is the secret spending key
#
#      %s
#
#  KEEP IT SAFE, IT HAS NOT BEEN SAVED ANYWHERE
#
#  To spend, import the secret key with
#
#      zcash-cli z_importkey KEY rescan
#
###############################################################
#
#  The following is a mnemonic encoding of the secret key
#  which you can write down as a paper wallet
#      
#      %s
#      %s
#      %s
#
#  Run "zcash-mini -mnemonic" to rebuild your secret key
#
###############################################################
#
#  Finally, here is the viewing key
#
#      %s
#
#  (not yet supported by the full node)
#
###############################################################

`

func main() {
	simpleMode := flag.Bool("simple", false, "output only values without decoration or text")
	existingKey := flag.Bool("key", false, "ask for an existing spending key instead of generating one")
	mnemonicKey := flag.Bool("mnemonic", false, "rebuild the key from a mnemonic phrase")
	vanityPrefix := flag.String("prefix", "", "search for an address with a given prefix")
	vanityRegexp := flag.String("regexp", "", "search for an address matching a given regexp - SLOW")
	flag.Parse()

	var rawKey []byte
	switch {
	case *existingKey:
		rawKey = readKey()
	case *mnemonicKey:
		rawKey = readMnemonic()
	case *vanityPrefix != "":
		rawKey = zcash.GenerateVanityKey(*vanityPrefix, zcash.ProdAddress)
	case *vanityRegexp != "":
		rawKey = GenerateVanityKeyRegexp(*vanityRegexp)
	default:
		rawKey = zcash.GenerateKey()
	}

	rawAddr, err := zcash.KeyToAddress(rawKey)
	if err != nil {
		fatal(err)
	}
	rawViewKey, err := zcash.KeyToViewingKey(rawKey)
	if err != nil {
		fatal(err)
	}

	key := zcash.Base58Encode(rawKey, zcash.ProdSpendingKey)
	addr := zcash.Base58Encode(rawAddr, zcash.ProdAddress)
	viewKey := zcash.Base58Encode(rawViewKey, zcash.ProdViewingKey)

	words := bip39.Encode(rawKey)

	if *simpleMode {
		mnemonic := strings.Join(words, " ")
		fmt.Printf("%s\n%s\n%s\n%s\n", addr, key, mnemonic, viewKey)
	} else {
		w1 := strings.Join(words[:8], " ")
		w2 := strings.Join(words[8:16], " ")
		w3 := strings.Join(words[16:], " ")
		fmt.Printf(template, logo, addr, key, w1, w2, w3, viewKey)
	}
}

func readKey() []byte {
	fmt.Fprint(os.Stderr, `Enter a spending key ("SK" prefix): `)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	key := scanner.Text()
	rawKey, version, err := zcash.Base58Decode(key)
	if err != nil {
		fatal(err)
	}
	if version != zcash.ProdSpendingKey {
		fatal("this is not a spending key.")
	}
	return rawKey
}

func readMnemonic() []byte {
	fmt.Fprint(os.Stderr, "Enter the 24 words, separated by spaces: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	words := strings.Split(scanner.Text(), " ")
	if len(words) != 24 {
		fatal("a mnemonic must be 24 words long")
	}
	rawKey, corrections, err := bip39.Decode(words)
	if err != nil {
		fatal(err)
	}
	for _, c := range corrections {
		fmt.Fprintln(os.Stderr, "[INFO] Automatically corrected:", c)
	}
	return rawKey
}

func GenerateVanityKeyRegexp(vanityRegexp string) []byte {
	r, err := regexp.Compile(vanityRegexp)
	if err != nil {
		fatal("failed to compile the regular expression:", err)
	}

	for {
		rawKey := zcash.GenerateKey()
		rawAddr, err := zcash.KeyToAddress(rawKey)
		if err != nil {
			fatal(err)
		}
		addr := zcash.Base58Encode(rawAddr, zcash.ProdAddress)

		if r.MatchString(addr) {
			return rawKey
		}
	}
}

func fatal(v ...interface{}) {
	v = append([]interface{}{"[FATAL] Error:"}, v...)
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}
