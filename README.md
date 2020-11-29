go install
sudo vim ~/.bashrc
export PATH=$PATH:$GOPATH/bin
:wq! (save the file)
source ~/.bashrc
clone repository
add dependencies using go get command

run using
go run main.go