# GOROOT
goroot="/usr/local/go"
if [ -d "$goroot" ]; then
    export GOROOT="$goroot"
fi
# PJROOT
export PJROOT=~/ffp



# GOPATH
export GOPATH=$PJROOT/src
# PATH
export PATH=$PJROOT/bin:$GOROOT/bin:/bin:/sbin:/usr/sbin:/usr/bin:/usr/local/bin
