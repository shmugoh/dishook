#/bin/sh

ostype=`uname`
arch=`uname -m`
binary=dishook_$ostype\_$arch

if [ "$ostype" == "Darwin" ]; then
    cd /usr/local/bin
else
    cd /usr/bin
fi

curl -L -O https://github.com/juanpisuribe13/dishook/releases/latest/download/$binary
mv $binary dishook
chmod +x dishook
cd -;