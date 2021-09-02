# zcash-mini

`zcash-mini` is a minimal, portable Zcash wallet generator in Go.

```
$ go get -u github.com/FiloSottile/zcash-mini

 - or -

$ git clone https://github.com/FiloSottile/zcash-mini
$ cd zcash-mini
$ go mod init zcash-mini
$ go mod tidy
$ go build -mod=mod
$ sudo install ./zcash-mini /usr/local/bin
```

```
$ zcash-mini

###############################################################
#
#  Here is your z-address
#
#      zchb1pjPZj5km3arxocST98jY27BzFqiaK2f7vLgyYgStPSuQ1dVR97ahfbz51oQM3Xb8VooGh9E5dyfMN2SJ1q1HVcsExT
#
###############################################################
#
#  Here is the secret spending key
#
#      SKxtAQQL74P5HMN73niHX1YwYZbjBMBPzp8NQ2M35Z2TybUbjiKc
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
#      armed fortune seek athlete humor please margin prosper
#      spend stool weapon buzz verify radio hamster couple
#      exercise idea stock year elder pass dune aspect
#
#  Run "zcash-mini -mnemonic" to rebuild your secret key
#
###############################################################
#
#  Finally, here is the viewing key
#
#      112TXSCh37UifeAJMmf7jcDpRiGp7krSggfqRVWqSYQNwrL8wj2Y
#
#  (not yet supported by the full node)
#
###############################################################
```

To re-process an existing key instead of generating a new one, use `-key`.

To generate vanity addresses use `-prefix` or the very, very slow `-regexp`. There is no GPU support, so you won't get the performance you would get with [other implementations](https://github.com/plutomonkey/zcash-vanity) (which I have not used or reviewed).

To get script-friendly output use `-simple`.

To cross compile simply run e.g. `GOOS=linux GOARCH=arm make`.

**This is experimental software and it will eat your money, your hard drive and your pets.**

## Use cases

* offline and paper wallets
* non-Linux amd64 environments (Raspberry Pi, OpenBSD, ...)
* machines without the resources to run a full node

## Features

* offline generation of z-address wallets
* vanity addresses
* paper wallet / mnemonic generation
* 300 LoC, easy to review
* pure Go, easy to cross-compile / deploy

Balance management and spending operations are not supported.

## Intended workflow

* generate a z-address with `zcash-mini` on a secure machine
* optionally write down the 24 words as a paper wallet
* receive Zcash on a t-address in a full node
* send the Zcash to the z-address yourself
* to spend the Zcash, export the private key to a full node
