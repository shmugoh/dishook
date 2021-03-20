import os
import sys
import ctypes
from platform import system as platform

if platform() == 'Windows':
    # Windows Installation

    isAdmin = ctypes.windll.shell32.IsUserAnAdmin()
    if isAdmin:
        # Has admin privilages
        print("Installing...")
        os.popen('copy dishook.exe %WINDIR%')

        print("Installed! To run it, run dishook on your terminal.")
        print("If you want to uninstall it, please refer to %s." % (os.environ['WINDIR']))
        os.system('pause')
    else: 
        # No admin privilages
        print("No admin privilages detected. ")
        ctypes.windll.shell32.ShellExecuteW(None, "runas", sys.executable, " ".join(sys.argv), None, 1)

else: # Unix Installation
    
    print("Installing...")
    os.system('sudo cp dishook /usr/bin')
    print("Installed! To run it, run dishook on your terminal.")
    print("If you want to uninstall it, please refer to /usr/bin")

# the way these messages are written is temporal. will change it whenever i can