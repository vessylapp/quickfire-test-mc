# Vessyl QuickFire Test - Minecraft Server

## Example Run Command

`docker run --restart always -d -p 25565:25565 -p 25575:25575 -e Version=1.18.2 -e MOTD="My Minecraft Server" -e Rcon=true -e RconPassword="Password" -e MaxPlayers=5 image_name:tag`

## Description

This is a simple Minecraft server that runs on a Vessyl Instance via QuickFire. It is a vanilla server with no mods.

## How to use?

Set environment variables in the Vessyl dashboard and deploy the server. The server will be up and running in a few minutes.

## Environment Variables

- `Version`: Set this to the version of Minecraft you want to run.
- `MOTD`: Set this to the message of the day you want to display. Default is `Welcome to the Vessyl Minecraft Server!`.
- `Port`: Set this to the port you want to run the server on. Default is `25565`.
- `Seed`: Set this to the seed you want to use for the world. Default is `0`.
- `WhiteList`: Set this to `true` if you want to enable the whitelist. Default is `false`.
- `Hardcore`: Set this to `true` if you want to enable hardcore mode. Default is `false`.
- `Rcon`: Set this to `true` if you want to enable RCON. Default is `false`.
- `RconPassword`: Set this to the password you want to use for RCON. Default is `password`.
- `RconPort`: Set this to the port you want to use for RCON. Default is `25575`.
- `MaxPlayers`: Set this to the maximum number of players you want to allow. Default is `20`.
