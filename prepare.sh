#!/bin/bash

set eux;

WORKDIR=$(pwd)

echo "-------------------------------------------------------";
echo "Installing Dependencies";
echo "-------------------------------------------------------";

# Update mirrors
apt-get update;

# Install dependencies
apt install golang ca-certificates curl gnupg lsb-release -y;

# Download docker repository key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg;

# Add docker repository
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null;

# Update mirrors
apt-get update;

# Install docker for postgres
apt-get install docker-ce docker-ce-cli containerd.io -y;

if systemctl is-active --quiet postgresql.service; then
		systemctl disable --now postgresql;
fi

# Start database container
docker run --name postgres -e POSTGRES_PASSWORD=fim -e POSTGRES_USER=fim -p 5432:5432 -d postgres;

echo "-------------------------------------------------------";
echo "Compiling Server";
echo "-------------------------------------------------------";

cd "${WORKDIR}/fimserver" && go build;
chmod 0777 fimserver;

echo "-------------------------------------------------------";
echo "Compiling Client";
echo "-------------------------------------------------------";

cd "${WORKDIR}/fimclient" && go build;
chmod 0777 fimclient;

echo "-------------------------------------------------------";
echo "Compiling Agent";
echo "-------------------------------------------------------";

cd "${WORKDIR}/fimagent" && go build;
chmod 0700 fimagent;

echo "-------------------------------------------------------";
echo "Compiling Ransomware";
echo "-------------------------------------------------------";

cd "${WORKDIR}/ransomware" && go build;
chmod 0777 ransomware;

echo "-------------------------------------------------------";
echo "Preparing Database";
echo "-------------------------------------------------------";

cd "${WORKDIR}/fimserver" && ./fimserver --setup;

echo "-------------------------------------------------------";
echo "Preparation finished";
echo "-------------------------------------------------------";