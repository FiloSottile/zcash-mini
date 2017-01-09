package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

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
#  Here is your new z-address
#
#      %s
#
#  and here is the secret key
#
#      %s
#
#  and here is the viewing key (not yet supported by the full node)
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

`

func main() {
	simpleMode := flag.Bool("simple", false, "output only address and key")
	existingKey := flag.Bool("key", false, "ask for an existing spending key instead of generating one")
	vanityPrefix := flag.String("prefix", "", "search for an address with a given prefix")
	vanityRegexp := flag.String("regexp", "", "search for an address matching a given regexp - SLOW")
	flag.Parse()

	var rawKey []byte
	switch {
	case *existingKey:
		rawKey = readKey()
	case *vanityPrefix != "":
		rawKey = zcash.GenerateVanityKey(*vanityPrefix, zcash.ProdAddress)
	case *vanityRegexp != "":
		rawKey = GenerateVanityKeyRegexp(*vanityRegexp)
	default:
		rawKey = zcash.GenerateKey()
	}

	rawAddr, err := zcash.KeyToAddress(rawKey)
	if err != nil {
		log.Fatal(err)
	}
	rawViewKey, err := zcash.KeyToViewingKey(rawKey)
	if err != nil {
		log.Fatal(err)
	}

	key := zcash.Base58Encode(rawKey, zcash.ProdSpendingKey)
	addr := zcash.Base58Encode(rawAddr, zcash.ProdAddress)
	viewKey := zcash.Base58Encode(rawViewKey, zcash.ProdViewingKey)

	if *simpleMode {
		fmt.Printf("%s\n%s\n", addr, key)
	} else {
		fmt.Printf(template, logo, addr, key, viewKey)
	}
}

func readKey() []byte {
	fmt.Print(`Enter a spending key ("SK" prefix): `)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	key := scanner.Text()
	rawKey, version, err := zcash.Base58Decode(key)
	if err != nil {
		log.Fatal(err)
	}
	if version != zcash.ProdSpendingKey {
		log.Fatal("This is not a spending key.")
	}
	return rawKey
}

func GenerateVanityKeyRegexp(vanityRegexp string) []byte {
	r, err := regexp.Compile(vanityRegexp)
	if err != nil {
		log.Fatal("Failed to compile the regular expression:", err)
	}

	for {
		rawKey := zcash.GenerateKey()
		rawAddr, err := zcash.KeyToAddress(rawKey)
		if err != nil {
			log.Fatal(err)
		}
		addr := zcash.Base58Encode(rawAddr, zcash.ProdAddress)

		if r.MatchString(addr) {
			return rawKey
		}
	}
}
