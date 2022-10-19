#!/bin/bash
sudo apt update && sudo apt upgrade
sudo apt install postgresql postgresql-contrib

# WSL:
# sudo service postgresql status
# sudo service postgresql start
# sudo service postgresql stop

# Linux:
# sudo systemctl status postgresql
# sudo systemctl start postgresql
# sudo systemctl stop postgresql

# PostgreSQL create user:
# sudo -u postgres psql
# CREATE USER username PASSWORD 'password' CREATEDB;

# WSL - To connect to PostgreSQL on WSL2 from a Windows host
# you can use with any GUI you prefer, for example with pgAdmin
