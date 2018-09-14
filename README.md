# Go-React-Boilerplate
Boilerplate for SSAD Assignment 2

If you're using this repo as reference, star it.

# Go

## Setting up
- For Ubuntu: [Look here](https://www.linode.com/docs/development/go/install-go-on-ubuntu/)
- For MacOS: [Look here](http://sourabhbajaj.com/mac-setup/Go/README.html)
- Optional: If you want to have your executables to be stored in `bin`, add the following to your .bashrc:
```bash
              export GOBIN=$GOPATH/bin
```
- Installing third party packages; use `go get`. An example:-
```bash
        go get -u -v gopkg.in/mgo.v2
```
- You'll have to install the packages before you run the code, since my OS may be different from yours. Packages used in this tutorial -
```bash
              go get -u -v github.com/gin-gonic/gin
              go get -u -v github.com/jinzhu/gorm
              go get -u -v github.com/jinzhu/gorm/dialects/sqlite
              go get -u -v github.com/gin-contrib/cors
```
- To run a program:-
```bash
              go run file.go
```


# React

## Setting up
- First, install node. Then install yarn (because npm is not very nice)
- Ubuntu:
```bash
              curl -sL https://deb.nodesource.com/setup_10.x | sudo -E bash -
              sudo apt-get install -y nodejs
              npm install -g yarn
              yarn global add create-react-app
```
- MacOS:
```bash
              brew install node
              npm install -g yarn
              yarn global add create-react-app
```
- Then, create a new React application by running the following:
```bash
              create-react-app name_of_app
```
- In the app I've made, make sure you run ```bash npm install``` or just ```bash yarn``` inside the `react-app` folder. This command automatically looks at the `package.json` file and installs the required packages into the `node_modules` folder.
- You can run the app by running `yarn start`
