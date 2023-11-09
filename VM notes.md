# VM Setup Notes

## Install Virtual Box

<https://www.virtualbox.org/>

## Install Kubuntu

First Download it: <https://kubuntu.org/>

Use the following settings:

- computer name: vm
- username: vm
- password: ?
- I used a virtual sized hard drive of 500GB, Dynamically allocated
- 10GB of RAM
- set it to 4 processors
- set video ram to max. 128M

### Install Guest Module

Go to the top menu and select insert guest module CD.

Tried to install the guest module, but it still doesn't work.

To get the guest package to install, it said "please install the gcc make perl packages"
To do this, I tried the following command and it seemed to work. "sudo apt install gcc make perl"
tried to install the guest module again. It seems like it worked. said to restart.
restarted.
Working.

### Shared Folder with Host Computer

Now, to add a shared folder between the host PC and the VM. Go to the virtual box options and select the shared folders tab. Select adding a new folder.

- Path: `C:\Users\StandardUser\VirtualBox VMs\Kubuntu1\SharedFolderWithVM`
- Folder Name: `SharedFolderWithVM`
- Auto Mount: `true`
- Make Permanent: `true`

The file showed up at `/media` but I can't seem to access it.

Looking it up, I found this: <https://superuser.com/questions/307853/permission-denied-when-accessing-virtualbox-shared-folder-when-member-of-the-vbo>

So, if I add my user to a user-group vboxsf then it works.

```bash
sudo usermod -aG vboxsf $(whoami)
```

May need to reboot for it to take effect.

## Terminal Command History Search With PageUp/Down

use this command to enable history search:

```bash
code /etc/inputrc
```

Then uncomment the two lines as seen below.

```bash
# alternate mappings for "page up" and "page down" to search the history
"\e[5~": history-search-backward
"\e[6~": history-search-forward
```

An Alternate is to put the lines in `~/.inputrc` as suggested in <https://stackoverflow.com/questions/60153457/how-to-enable-history-search-by-page-up-down-in-git-bash-like-in-linux>
Basically, edit the `~/.inputrc` file. If it doesn't exist yet, that is okay.

```bash
code ~/.inputrc
```

Then add the following lines to the file.

```file
    "\e[5~": history-search-backward
    "\e[6~": history-search-forward
```

## Updating your VM

```bash
sudo apt update && sudo apt upgrade
```

## Notes on getting VirtualBox Out of Turtle Mode

- If you have installed WSL, then uninstall it in `apps & Features`
- with the start menu search: open `Turn Windows features on or off` and make sure the following are not unchecked
  - Hyper-V
  - Virtual Machine Platform
  - Windows Hypervisor Platform
  - Windows Subsystem for Linux

I have read that if `Memory integrity` is enabled, it also blocks VirtualBox from running at full speed.
