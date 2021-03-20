import os
import sys
import ctypes

from platform import system as platform

try:
    import requests
except:
    import pip
    print("Module requests not found. Installing...")
    pip.main(['install'], 'requests')
    del pip
    import requests
    print("Installed requests module!")

def getFile() :
    rep = "https://api.github.com/repos/juanpisuribe13/Dishook/releases/latest"
    rep_json = requests.get(rep).json() 
    release_assets = len(rep_json["assets"])

    for i in range(release_assets):
        release_url = rep_json["assets"][i]['browser_download_url']

        if platform() == 'Windows' and release_url[-3:] == 'exe':
            break
        elif platform() != 'Windows' and release_url[-3:] != 'exe':
            break
        else:
            continue

    print("Downloading from %s..." % (release_url))
    release_raw = requests.get(release_url).content

    if release_url[-3:] == 'exe':
        open('dishook.exe', 'wb').write(release_raw)
        print("Downloaded!")
    else:
        open('dishook', 'wb').write(release_raw)
        print("Downloaded!")

if platform() == 'Windows':
    # Windows Installation
    isUserAdmin = ctypes.windll.shell32.IsUserAnAdmin()
    if isUserAdmin:
        # Has admin privilages
        getDownloadURL()
        print("Installing...")
        os.popen('copy dishook.exe %WINDIR%')

        print("Installed! To run it, run dishook on your terminal.")
        print(f"If you want to uninstall it, please refer to {os.environ['WINDIR']}")
        os.system('pause')
    else: 
        # No admin privilages
        print("To continue, you will need to provide admin privilages.")
        print(f"Dishook will be installed in {os.environ['WINDIR']} because PATH is already set up there.\n", 
        "(so you can use dishook without being in the same folder where the executable is).\n")
    
        os.system('pause')
        ctypes.windll.shell32.ShellExecuteW(None, "runas", sys.executable, " ".join(sys.argv), None, 1)

else: # Unix Installation
    getDownloadURL()
    print("Installing...")
    os.system('sudo cp dishook /usr/bin')
    print("Installed! To run it, run dishook on your terminal.")
    print("If you want to uninstall it, please refer to /usr/bin")

# the way these messages are written is temporal. will change it whenever i can