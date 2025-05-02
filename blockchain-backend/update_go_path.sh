#!/bin/bash

# Append Go to the PATH in ~/.zshrc
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc

# Reload ~/.zshrc to apply changes
source ~/.zshrc

echo "Go path has been added to ~/.zshrc and reloaded."

