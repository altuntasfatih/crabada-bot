# Crabada Looter Bot

[Crabada](https://www.crabada.com/) is a fully decentralized play-and-earn idle game.
Looting is a mission in the game, and it is more profitable than others. But there was an intense competition between player and bots to loot open Mines, So manuel looting was almost impossible. Therefore, we developed this bot to loot open mines very quickly. It interacts directly with Crabada Smart Contract.

It worked very successfully until anti bot patch. Now there is recaptcha to block bots, and it seems working well. We decided to share this repo with public.
We hope that it can be useful someone who curious.

## How is it working ?

There are two different looting mode.

### Subscribing Mode

Init phase:
* It gets your available team's info from CrabadaApi or smart contract(both of can be used).

Subscribing phase:

1. Listen StartGame(opened mine) events from smart contract.
2. Look Mine's team info from cache if cache is enabled(older versions gets team's info from smart contract, but it
   caused extra i/o and slowness)
3. If looting is possible(it means your team is stronger than miner's team), it call attack function of crabada smart
   contract.
4. Continue this process until team is locked with looting mission.After context is cancelled all goroutines exits and stop script.

### Iterator Mode

We realized GameId(or mineId) is incremental number. When a player opens a mine, newGameId is always one more of last
gameId. To get ahead of other bots, we decided to loot directly game Id without waiting StartGame Event(It is like
a brute force attack).

Init phase:

* It gets your available team info from CrabadaApi or smart contract(both of can be used).
* Get last gameId from CrabadaApi.

Iteration phase:

1. Increase gameId by one.
2. Get GameInfo(gameId) from contract, if game is available, continue next step.
3. Look Mine's(Game) team info from cache if cache is enabled.
4. If it is looting possible, it calls attack function of smart contract.
5. Continue this process until team is locked with looting mission. After context is cancelled all goroutines exits and stop script.

### Summary
It was a very excitement process for us. We have tried many things to beat others bots.
A few of them :

* Using custom Avalance node to speed up Transactions.
* To loot only players who usually doesn't reinforce.
* Tuning gasPrice, gasTip and GasFeeCap(LegacyTx or eip-1559) to speed up Transactions.
* Changing server geolocation to decrease network latency.
* Subscribing multiple node at same time to listen StarGame event.
* Management transactions' nonce by bot.
* There is a load balancer in front of api.avax.network. It blocks us when we request too many. To bypass it, we splited contract three diffrent instance: CallerContract, TransactorContract and FiltererContract. Each of them are using diffrent node client so that way we reduced request traffic which is going to one node.
* To limit cost of transactions' fee, we added different thresholds to stop script.

Not: We found Crabada contract abi inside of javascript files(crabada.com)

### Install

`go install`

### Usage

`crabada-bot -h`

#### Cache Reload

`crabada-bot reload -h`

`crabada-bot reload -c 1`

#### Subscription mode

`crabada-bot loot -h`

`crabada-bot loot --key <privateKey> --gasTip <GasTip> --cache`

#### Iterator mode

`crabada-bot loot-iterate -h`

`crabada-bot loot-iterate --key <privateKey> --gasTip <GasTip> --cache`