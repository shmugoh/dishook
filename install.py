import os
import sys
import ctypes
from platform import system as platform

def is_admin():
    try:
        return ctypes.windll.shell32.IsUserAnAdmin()
    except:
        return False

if platform() == 'Windows':
    if is_admin():
        dst = os.environ['WINDIR']
        os.popen('copy dishook.exe %WINDIR%')
        os.system('pause')
    else:
        print(sys.argv)
        ctypes.windll.shell32.ShellExecuteW(None, "runas", sys.executable, " ".join(sys.argv), None, 1)
else: # unix
    os.system('sudo cp dishook /usr/bin')