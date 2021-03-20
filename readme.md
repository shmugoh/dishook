# Dishook (Beta)
Are you tired of setting up a lot of stuff you don't even know about just to mess up with Discord's webhook? Or even better, are you tired wasting time because of going to that one file you have to run your webhook?

Well, now you don't need to worry about that!

Dishook can send your messages thru webhook on a command-line! Just give it the webhook link, 
the message you wish and profit!

It saves you a lot of time. Instead of downloading a programming language, installing the correct libraries, 
copying and pasting code, then trying to fix it, you could just use this instead.

For the nerds: This is written in Golang. Installer is written in Python.

And please, don't whine about the installer being written in **THAT** language. ~(i see you rand :eyes:)~

# Installation instructions (recommended)

You can use dishook without "installing" it while being in the same folder. But if you don't want to
switch folders a lot just to send a message, do this.

## Requirements
- Python 3 if you're gonna do automatic installation. If you're on Windows, I recommend you to get it 
[on Python's official website](https://www.python.org/downloads/) instead of the one from the Microsoft Store.

- Common Sense. Please, I don't want some kid to submit an issue asking what's their computer's administrator password 
to access C:\Windows.

- That's it.

## Automatic Installation

**1.** Download the Installer [here](https://raw.githubusercontent.com/juanpisuribe13/Dishook/main/install.py)

**2.** Go to a command-line and run `python3 install.py` (or `python install.py` if python3 opens up Microsoft Store).

**2.1** If the requests module isn't found, the script will install it automatically. The request module helps us 
connect to the internetz.

**3.** Done! 

**For Windows Users:** If you're on Windows though, your antivirus may detect it as a threat (because of it being 
saved on C:\Windows). Add an exclusion on it. 
If you're dumb enough to think it's a virus, do me a favour and read the source code.

## Manual Installation

**1.** Go to the repository's [latest release](https://github.com/juanpisuribe13/Dishook/releases/latest).

**2.** Download the executable. If you have a hard time (lol) telling which one is for your operating system, 
.exe is for Windows and the one without a file type is for Linux.

**3.** Move the executable to your operating system's path. For Windows, it is C:\Windows. For Linux, /usr/bin.

**4.** Done!

# FAQ

### How do I use this?

Open a new command-line window, and type in like this:
```bash
dishook https://discord.com/api/webhooks/.../.../ Hello World!
```

### As any other command-line program, will arguments be implemented!

Yep! If you wanna know which arguments will be implemented, refer to Projects.

### Why is the installer written in Python instead of Golang?

It's something personal. I wanna test having different programming languages in a single project.

# Contributing

If you see that the code's messy, something doesn't work, or I left something dumb there accidentally then feel 
free to send a pull request and I'll review it!

# License
[MIT](https://raw.githubusercontent.com/juanpisuribe13/Dishook/main/LICENSE)
