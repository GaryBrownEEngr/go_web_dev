# Readme

## Basic Linux Setup

See the file `VM notes.md`

## Git and GitHub

```bash
sudo apt-get install git
```

Now we need to setup an SSH key with GitHub:

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```

More info can be found [here](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent).

when prompted to enter where to save the file, just press enter to use the default. Then press enter again to use no passphrase.  
Now we need to copy the public key to GitHub. Print out the key using:

```bash
cat ~/.ssh/id_ed25519.pub
```

Then in GitHub, Select: `Settings > SSH and GPG keys > New SSH key`  
more info can be found [here](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account?platform=linux).
Also enter the key again in GitHub as a signing key.

Now, we should be able to clone the repo using the SSH link.

```bash
git clone git@github.com:GaryBrownEEngr/go_web_dev.git
```

### Private repos and Golang

If you have private repos that golang will need to build from. You will need to configure a replace directive inside of git. Then git will swich to using SSH mode, and use the key, to fetch the repos.

```bash
git config --global --add url."git@github.com".insteadOf "https://github.com"
```

Now we setup automatic signing of commits

```bash
git config --global gpg.format ssh
git config --global user.signingkey ~/.ssh/id_ed25519.pub
git config --global commit.gpgsign true
```

Now, whenever to you create a commit, it will be signed with your key and marked as verified on GitHub.

### Set Name and Email

In your GitHub `Settings > Email` you should check the `Keep my email addresses private` setting.  
Then, with git set your email and username:

```bash
git config --global user.email "ID+USERNAME@users.noreply.github.com"
git config --global user.name "NAME"
```

Then, use the following commands to show all the git configurations you have set.

```bash
cat ~/.gitconfig
```

More info [here](https://docs.github.com/en/account-and-profile/setting-up-and-managing-your-personal-account-on-github/managing-email-preferences/setting-your-commit-email-address).
  
## Install VS Code

To install vscode, <https://code.visualstudio.com/docs/setup/linux> says to use the command: `sudo apt install ./<file>.deb`
Then to run it, just type in "code"

### Suggested VS Code Extensions

- Docker
- ESLint
- GitLens
- Go
- markdownlint
- Prettier - Code formatter

I usually turn on auto save in `File > Auto Save`

## Golang

### Install Go

The instructions can be found [here](https://go.dev/doc/install).

```bash
wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
```

Now we need to set the env variables.

```bash
code ~/.bashrc
```

Add the following to the bottom:

```bash
# Golang
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$(go env GOPATH)/bin
# as talked about here: https://go.dev/doc/gopath_code#GOPATH
# the /bin needs to added to the path to be able to run "go install"ed things.

# helpful shortcuts:
alias bashrc="code ~/.bashrc"
alias src="source ~/.bashrc"

alias gitgraph="git log --all --decorate --oneline --graph"
alias gitfetch="git fetch --prune --prune-tags"

alias githubactionsgo="echo tidy && go mod tidy && echo build && go build ./... && echo vet && go vet ./... && echo test && go test ./... && echo lint && golangci-lint run"
alias gocoveragehtml="go test -short ./... -coverprofile coverage.out && go tool cover -html=coverage.out -o coverage.html && sleep 2 && firefox coverage.html"
```

Now either source the update to your terminal, or restart the terminal:

```bash
source ~/.bashrc
```

Now verify your installation.

```bash
go version
```

### install golangci-lint

This uses the repo file: `.golangci.yaml`.  
Go [here](https://golangci-lint.run/usage/install/) for full instructions. The easiest option is to go install it with the following command:

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.0
```

Then open your VS code settings and search for `go lint` and update the option `Go: Lint Tool` to be set to `golangci-lint`

### install gofumpt

<https://github.com/mvdan/gofumpt>

```bash
go install mvdan.cc/gofumpt@latest
```

## Docker

### Install Docker

Go [here](ttps://docs.docker.com/engine/install/ubuntu/) for full instructions.

```bash
sudo apt-get remove docker docker-engine docker.io containerd runc
sudo apt-get update
sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

sudo mkdir -m 0755 -p /etc/apt/keyrings

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

  sudo apt-get update
  sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

### Make docker work without sudo

Add your user to the docker group using the command:

```bash
sudo groupadd docker
sudo usermod -aG docker ${USER}
```

Then log out and log back in.

Then test if it worked with `docker run hello-world`

## Random Tips

If you want to remove all the dangling docker images: <https://nickjanetakis.com/blog/docker-tip-31-how-to-remove-dangling-docker-images>

```bash
docker rmi -f $(docker images -f "dangling=true" -q)
```

remove all non-running images:

```bash
docker rm $(docker ps -aq)
```

### Test hello world

```bash
docker run --name hello-world hello-world
# should download, build, and start an example hello world image.

# now to show it and its status
docker ps -a

# now remove the image
docker rm hello-world
```

## Test locally

```bash
docker build -t twertle .
docker run -p 10000:10000 twertle

docker ps # show running containers
docker ps -a # show all containers
docker exec -it loving_hamilton /bin/sh # start a shell session inside the container
docker rm loving_hamilton # delete a docker image.
```

<http://localhost:10000/>  
<http://localhost:10000/guess.html>  
<http://localhost:10000/ticktacktoe>  
<http://localhost:10000/test2.html>  
<http://localhost:10000/api/articles>

## Keep Windows From Locking

Open `Windows PowerShell` In the start menu and enter the following script:

```bash
$WShell = New-Object -com "WScript.shell"
while($true){
  $WShell.sendkeys("{SCROLLLOCK}")
  start-sleep -milliseconds 100
  $WShell.sendkeys("{SCROLLLOCK}")
  start-sleep -seconds 240
}
```
