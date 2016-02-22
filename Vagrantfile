# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrant configuration to do these things:
# 1. Spin up a Ubuntu 14 instance
# 2. Install golang with $GOHOME set to /home/vagrant/gohome
# 3. Sync folder ./merkle -> $GOHOME/merkle-tree
Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu/trusty64"
  config.vm.synced_folder "euler", "/home/vagrant/gohome/src/euler"

  config.vm.provision "shell", path: "deploy/install_utils.sh"
  config.vm.provision "shell", path: "deploy/install_go.sh", privileged: false
  config.vm.provision "shell", path: "deploy/source_bashrc.sh", privileged: false
end
