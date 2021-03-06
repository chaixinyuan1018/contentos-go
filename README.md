# Contentos-go
[![Build Status](https://travis-ci.com/coschain/contentos-go.svg?branch=master)](https://travis-ci.com/coschain/contentos-go)

Official golang impementation of the Contentos protocol

Contentos Website: https://www.contentos.io

Contentos White Paper: https://www.contentos.io/subject/home/pdfs/white_paper_en.pdf

Follow us on https://twitter.com/contentosio

Join discussion at https://t.me/ContentoOfficialGroup

**WARNING:** The branch is under heavy development. Breaking changes are actively added.

**Note**: Requires [Go 1.11.4+](https://golang.org/dl/)

## Building the source

pull source code
```bash
git clone git@github.com:coschain/contentos-go.git
```
build cosd and wallet
```
cd cmd/cosd
go build
cd ../wallet-cli/
go build
```

## Executables
The contento-go contains two executables as follow:

**cosd**: the daemon to run a local blockchain

**wallet**: the cli to interactive with chain.


## Running cosd

### Initialization

init will create a folder to hold cosd's running data,this will create a folder ~/.coschain/cosd
```bash
cosd init
```

you can also custom your own folder name, this will create a folder ~/.coschain/yourownname

```bash
cosd init -n yourownname
```

### Configuration
After initialization, configurations will be found under homedir_.cosd_nodename

The nodename is cosd or yourownname.

You can modify it if you like as long as you know what you are doing.

## Running multi nodes in single machine

if you want to set up a contentos network in single machine, you need to do:

1.run cosd init -n name to create different node folders,after call cosd init, there will be a output "Cosd running version:  defaultVersion", this mean cosd use default config,then edit all config.toml,make sure HTTPListen and RPCListen and NodeConsensusPort and NodePort and HealthCheck all unique to other node's config.toml.

```
./cosd init -n cos1
./cosd init -n cos2
./cosd init -n cos3
./cosd init -n cos4
vi ~/.coschain/cos1/config.toml
vi ~/.coschain/cos2/config.toml
vi ~/.coschain/cos3/config.toml
vi ~/.coschain/cos4/config.toml
```
![config1](doc/technical-whitepaper/assets/1.png)

![config2](doc/technical-whitepaper/assets/2.png)

![config3](doc/technical-whitepaper/assets/3.png)

2.start first cosd,use wallet connect cosd via RPC address according to config.toml.you can find privateKey_of_initminer in wallet_doc_cn.md(at wallet-cli folder)
```
go to cosd folder
./cosd start -n cos1
go to wallet-cli folder
./wallet-cli
switchport ip_to_your_node:8888
import initminer privateKey_of_initminer
```

3.create 3 new accounts and use info command to get thier public keys and priviate keys.
```
stake initminer initminer 1.000000
create initminer blockproducer1
create initminer blockproducer2
create initminer blockproducer3
info blockproducer1
info blockproducer2
info blockproducer3
```

4.edit other node's config.toml,change BootStrap to false,set LocalBpName and LocalBpPrivateKey that you
just created in step 3.

![config4](doc/technical-whitepaper/assets/4.png)

5.start remain cosd.
```
./cosd start -n cos2
./cosd start -n cos3
./cosd start -n cos4
```

6.use wallet again to regist 3 accounts as new bp
```
unlock iniminer
bp register blockproducer1 blockproducer1_publicKey
bp register blockproducer2 blockproducer2_publicKey
bp register blockproducer3 blockproducer3_publicKey
```

now job is finished.

optinal, the first cosd has been changed to a observe node since you regist bp, if you want first cosd also become a bp node, you need to create a new account,modify first cosd's config.toml,change BootStrap to false,set LocalBpName and LocalBpPrivateKey to new account's info(same as step 4),restart first cosd, regist new account as bp(same as step 6).

## Running multi nodes in multi machines

assume you have 4 machines and want to set up a contentos network, here the steps
you need to do:

1.choice a machine,run cosd init then edit config.toml,modify all ip relative items to your custom.
```
./cosd init
vi ~/.coschain/cosd/config.toml
```

2.start first cosd,use wallet connect cosd via RPC address according to config.toml.
```
switchport ip_to_your_node:8888
import initminer privateKey_of_initminer
```

3.create 3 new accounts and remember thier public keys and private keys.
```
stake initminer initminer 1.000000
create initminer blockproducer1
create initminer blockproducer2
create initminer blockproducer3
info blockproducer1
info blockproducer2
info blockproducer3
```

4.do cosd init on other 3 machines,edit config.toml,change BootStrap to false,set LocalBpName and LocalBpPrivateKey that you
just created in step 3.

5.start remain cosd on each machine.
```
./cosd start 
```

6.use wallet again to regist 3 accounts as new bp
```
unlock iniminer
bp register blockproducer1 blockproducer1_publicKey
bp register blockproducer2 blockproducer2_publicKey
bp register blockproducer3 blockproducer3_publicKey
```

now job is finished.

optinal, the first cosd has been changed to a observe node since you regist bp, if you want first cosd also become a bp node, you need to create a new account,modify first cosd's config.toml,change BootStrap to false,set LocalBpName and LocalBpPrivateKey to new account's info(same as step 4),restart first cosd, regist new account as bp(same as step 6).

### Interaction
enter

```bash
wallet
```

to get into interactive mode.

You can using some commands as below:

* account
* bp
* claim
* close
* create
* follow
* genKeyPair
* import
* info
* locked
* list
* load
* lock
* post
* reply
* transfer
* unlock
* transfer_vest
* vote

you can add `--help` or `help [command]` to get more detail infos.

## Contribution
Contributions are welcomed.

If you'd like to help out with the source code, please send a pull request. Or you can contact us directly by joining telegram.
