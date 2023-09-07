# PolitoGroupHelp_bot
A Telegram bot to help admins of a network telegram groups. Created on the trail of Group Help, it adds some functionalities that we, as a group, found missing.

## Code contribution instructions
The branch *main* contains the (stable and tested) code that can be safely be put into production.

For adding new features, please create a new branch *dev/<branch_name>*. Once you have cloned this repository, create a new branch from it, using git switch or git checkout -b. You can't directly push to *main*, so you can't destroy anything here.

Be sure to not publish any token or user ID in this repository, since it could be potentially seen from other malicious users.

## Features
*No features yet... add some by creating a new branch and a pull request*

## Software architechture
This bot is divided into microservices, that communicate with each other via RPC

Each folder in the main directory corresponds to a microservice, which are the following:

### telegramConnection
It is in charge of communicating towards Telegram APIs:
- it receives messages on behalf of the bot and forwards them to the other services for processing
- it receives orders from the other services to send requests to the Telegram API
