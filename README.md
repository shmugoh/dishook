# Dishook
Dishook is an efficient CLI Discord webhook executor written in Golang.

It helps save the user time - instead of downloading a programming language, installing the correct libraries, 
copying and pasting code, then trying to fix it, you could run this one simple ~~trick~~ command instead!

# Automatic Installation

## Windows (W.I.P)
**1.** Open PowerShell with administrator privilages.

**2.** Run `Get-ExecutionPolicy`. If it returns `Restricted`, run `Set-ExecutionPolicy AllSigned`.

**3.** 
```powershell
iwr -useb [SCRIPT_URL] | iex
```

## Linux

```bash
cd /usr/bin && { sudo curl -L -O https://github.com/juanpisuribe13/Dishook/releases/latest/download/dishook; sudo chmod +x dishook; cd -; }
```

# Manual Installation

**1.** Go to the repository's [latest release](https://github.com/juanpisuribe13/Dishook/releases/latest).

**2.** Download the executable. If you have a hard time telling which one is for your operating system, 
.exe is for Windows and the one without a file type is Linux.

**3.** Move the executable to your operating system's PATH.

**Windows:** Run `set PATH` and pick whichever folder you're comfortable putting it on.

**Linux:** /usr/bin

**4.** Done!

# FAQ

### Dishook is an invalid command, what do I do?

Check if dishook is already in your system's path. If it is, start a new terminal session and try again.

### How do I use this?

Open your terminal of choice and type in `dishook -h` to get started.

### There's a lot of names out there and you picked Dishook; there's a lot of projects with that name. Why?

Memory Muscle, and something like `discord-webhook` would look weird in terminal.

# Contributing

If you see that the code's messy, something doesn't work, or I left something dumb there accidentally then feel 
free to send a pull request and I'll review it!

# License
[Apache License 2.0](https://raw.githubusercontent.com/juanpisuribe13/Dishook/main/LICENSE)
