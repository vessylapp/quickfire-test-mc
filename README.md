# Vessyl QuickFire Test - Minecraft Server 1.21

## Description

This is a simple Minecraft server that runs on a Vessyl Instance via QuickFire. It is a vanilla server with no mods. The server is running on the latest version of Minecraft(1.21).

## How to use?

Set environment variables in the Vessyl dashboard and deploy the server. The server will be up and running in a few minutes.

## Environment Variables

- `MOTD`: Set this to the message of the day you want to display. Default is `Welcome to the Vessyl Minecraft Server!`.
- `MEMORY`: Set this to the amount of memory you want to allocate to the server. Default is `1G`.
- `PORT`: Set this to the port you want to run the server on. Default is `25565`.
- `SEED`: Set this to the seed you want to use for the world. Default is `0`.
- `GAMEMODE`: Set this to the gamemode you want to use. Default is `survival`.
- `DIFFICULTY`: Set this to the difficulty level you want to use. Default is `normal`.
- `WHITELIST`: Set this to `true` if you want to enable the whitelist. Default is `false`.
- `WHITELIST_PLAYERS`: Set this to a comma separated list of players you want to whitelist. Default is `''`.
- `OPS`: Set this to a comma separated list of players you want to op. Default is `''`.