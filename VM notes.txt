computer name: vm
username: vm
password: ?

I used a virtual sized hard drive of 500GB
8GB of RAM
set it to 4 processors
set video ram to max. 128M


Tried to install the guest module, but it still doesn't work. 

To get the guest package to install, it said "please install the gcc make perl packages"
To do this, I tried the following command and it seemed to work. "sudo apt install gcc make perl"
tried to install the guest module again. It seems like it worked. said to restart.
restarted.
Working.

Installed Docker using the following instructions.
https://docs.docker.com/engine/install/ubuntu/


Add your user to the docker group using the command: sudo groupadd docker
udo usermod -aG docker ${USER}
then log out and log back in.

docker run hello-world


To install vscode, https://code.visualstudio.com/docs/setup/linux says to use the command: "sudo apt install ./<file>.deb"
Then to run it, just type in "code"
added the Go addon and the "Prettier - Code formatter" addon


went to go's website, and followed instructions to download and install go.  https://go.dev/doc/install
To update the PATH variable, use the text editor nano to edit the file ~/.profile
"nano ~/.profile"
Then we added the following to the end of the file.
"
# Add golang to the path
PATH="$PATH:/usr/local/go/bin"
"
Had to log out and log back in for the path change to take effect.



Created snapshot 1



Trying to get github to work.
So, if we create an SSH key with the command "ssh-keygen -m PEM -t rsa -b 4096" and accept the standard location and use no password, then copy the public portion to github.

Then clone the git repo using the SSH link and it works.

Now I am trying to add a shared folder between the host PC and the VM. I set the docker project as the folder and set it as read only and auto mount. Also make perminent, not sure if we actually want this. The file showed up at /media but I can't seem to access it. 

Looking it up, I found this: https://superuser.com/questions/307853/permission-denied-when-accessing-virtualbox-shared-folder-when-member-of-the-vbo

So, if I add my user to a user-group vboxsf then it works.
sudo usermod -aG vboxsf $(whoami)
May need to reboot for it to take effect.



use this command to enable history search:
sudo nano /etc/inputrc

Then uncomment the two lines as seen below.
# alternate mappings for "page up" and "page down" to search the history
"\e[5~": history-search-backward
"\e[6~": history-search-forward



Docker:
remove all non-running images: "docker rm $(docker ps -aq)"


Created snapshot 2


to install NVM https://github.com/nvm-sh/nvm 
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash

Then you create the react app: "npx create-react-app abc"
Then you go into that directory: "cd frontend"
then start the app: "npm start"

