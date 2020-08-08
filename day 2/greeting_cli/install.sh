#!/bin/bash

# download binary file from GitHub
curl https://raw.githubusercontent.com/103cuong/greeting_cli/master/greeting_cli --output greeting_cli
# move this file to /usr/local/bin
sudo mv greeting_cli /usr/local/bin
# set permission for this cli
sudo chmod +x /usr/local/bin/greeting_cli
