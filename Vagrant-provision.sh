#!/bin/bash
# usage: Vagrant-provision.sh

set -eo pipefail

# deps
sudo apt-get update
echo "installing the essentials"
sudo apt-get install -y curl git mercurial make binutils bison gcc build-essential
echo "installing PostgreSQL"
sudo apt-get install -y postgresql postgresql-contrib
echo "installing redis"
sudo apt-get install -y redis-server

GO_VERSION=go1.4.2

if [ -d /home/vagrant/.gvm ]; then
  echo "Already installed gvm"
else
  echo "Installing gvm to manage Go versions"
  bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
fi

echo "installing Go ${GO_VERSION}"
[[ -s "/home/vagrant/.gvm/scripts/gvm" ]] && source "/home/vagrant/.gvm/scripts/gvm"
/home/vagrant/.gvm/bin/gvm install ${GO_VERSION}
gvm use ${GO_VERSION}

echo "installing Godep"
go get -u github.com/tools/godep
echo "installing forego"
go get -u github.com/ddollar/forego
echo "installing go-hotreload"
go get -u github.com/ivpusic/go-hotreload/hr

echo "populating user's bashrc"

add_gvm_to_path="export PATH=\$PATH:/home/vagrant/.gvm/bin/"
setup_go_command="gvm use ${GO_VERSION}"
default_dir_command="cd \$GOPATH/src/github.com/liveplant/liveplant-server"
fancy_terminal_setup="export PS1=\"\[\$(tput bold)\]\[\$(tput setaf 2)\][\u]\\$ \[\$(tput sgr0)\]\""
profile_file="/home/vagrant/.bashrc"

for command_string in "$add_gvm_to_path" "$setup_go_command" "$default_dir_command" "$fancy_terminal_setup"; do
  if ! grep -q "${command_string}" $profile_file; then
    echo "${command_string}" >> $profile_file
  fi
done

echo "symlinking project directory into GOPATH"
mkdir -p $GOPATH/src/github.com/liveplant
[[ -e $GOPATH/src/github.com/liveplant/liveplant-server ]] || ln -s /vagrant $GOPATH/src/github.com/liveplant/liveplant-server

echo "all done!"
