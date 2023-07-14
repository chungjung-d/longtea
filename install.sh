sudo apt update && sudo apt upgrade -y
sudo apt install -y git curl 

curl -o longtea https://github.com/chungjung-d/longtea/releases/download/v0.1.0/longtea_v0.1.0
curl -o longteac https://github.com/chungjung-d/longteac/releases/download/v0.1.0/longteac_v0.1.0

chmod +x longtea longteac
sudo mv longtea longteac /usr/local/bin

