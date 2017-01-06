# zcash-mini

`zcash-mini` is a minimal, portable Go Zcash wallet.

```
$ go get -u github.com/FiloSottile/zcash-mini
```

**This is experimental software and comes with no promises of not eating your money.**

## Use cases

* offline and paper wallets
* non-Linux amd64 environments (Raspberry Pi, OpenBSD, ...)
* machines without the resources to run a full node

## Features

* offline generation of z-address wallets
* paper backup generation
* private key export
* X LoC, easy to review
* pure Go, easy to cross-compile / deploy

Balance management and spending operations are not supported.

## Intended workflow

* generate a z-address with `zcash-mini` on a secure machine
* optionally turn it into a paper wallet
* receive Zcash on a t-address in a full node
* send the Zcash to the z-address yourself
* to spend the Zcash, export the private key to a full node
