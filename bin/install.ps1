$loc = "$env:APPDATA\dishook"
$win32_binary_url = "https://github.com/juanpisuribe13/Dishook/releases/latest/download/dishook.exe"

mkdir $loc
setx PATH "$env:path;$loc"

Start-BitsTransfer -Source $win32_binary_url -Destination $loc\dishook.exe

Write-Output "Installed! Please start a new terminal session"