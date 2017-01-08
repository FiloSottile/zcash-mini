# zcash-mini

`zcash-mini` is a minimal, portable Go Zcash wallet.

```
$ go get -u github.com/FiloSottile/zcash-mini
$ zcash-mini
###############################################################
#
#  Here is your new z-address
#
#      zcZvkjKo24G8LzSiU9UEvPeeu6rCYJc46bhe5dsYLqNgcR12GFCQsE6Z6w4LH3Mb82aYoGgpjRpK8VcwTesaFbpPZhmkCJe
#
#  and here is the secret key
#
#      SKxoocXqMjxsiE3DdEguWhcbqf3vz45rpMdmVqXPwzReBBDhXj6v
#
#  KEEP IT SAFE, IT HAS NOT BEEN SAVED ANYWHERE
#
#  To spend, import it with
#
#      zcash-cli z_importkey KEY rescan
#
###############################################################
```

**This is experimental software and it will eat your money, your hard drive and your pets.**

## Use cases

* offline and paper wallets
* non-Linux amd64 environments (Raspberry Pi, OpenBSD, ...)
* machines without the resources to run a full node

## Features

* offline generation of z-address wallets
* paper backup generation (TODO)
* 150 LoC, easy to review
* pure Go, easy to cross-compile / deploy

Balance management and spending operations are not supported.

## Intended workflow

* generate a z-address with `zcash-mini` on a secure machine
* optionally turn it into a paper wallet
* receive Zcash on a t-address in a full node
* send the Zcash to the z-address yourself
* to spend the Zcash, export the private key to a full node
