> [!WARNING]
> This is a **NSFW** chat system, oriented on adults.

# "The Gag Roulette" and "The Gag Curse" system

This is a system of Custom Commands for [YAGPDB.xyz Discord Bot](https://yagpdb.xyz/), implementing a BDSM-themed textual gagging mechanics on a Discord server.

## Features

* Participating users can "receive" a random gag from the bot, and get a 1 minite long timeout (if the user has consented beforehand).
* In a channel dedicated for **the Gag Roulette**, any sent message (with high probability) will "reward" user with a random gag. There are several gag categories and plenty of gag kinds, all with uneven probabilities to receive.
* Users with **the Gag Curse** active are having a small chance to become "gagged" after sending any message anywhere on the server. The curse can be set up with `curse` command, if the user has a NSFW access role.
* Submissive people may consent to be muted (timeouted) for 1 minute each time after receiving a gag (to feel themselves "gagged" for real), by aquiring a certain role.

## Discord server setup

* Add a NSFW access role (`SPICY`).
* Add a "gag curse" role (`Gag Curse`).
* Add a "consented to be timeouted role" (`Consent to Mute`).
* Add a channel for the "roulette" game (`😷the-gag-roulette`), available for users with the NSFW role.
* Add YAGPDB.xyz bot.
  * Enable timeouts.
  * Import these Custom Commands and tweak them for one's needs.
