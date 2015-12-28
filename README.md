# go-useful-scripts
Some useful scripts that would be done usually in bash

#### `exportToBashrc`
Will insert a `export` in your home .bashrc file.
* Takes 2 arguments: [variable_name] and [path]

##### Example
```bash
exportToBashrc gopath /home/sayden/go
```
Will insert `export GOPATH=/home/sayden/go` at the end of your .bashrc file (uppercasing first word)


#### `adLocalGitUser`
Will add the user.name and user.email properties to override your global Git user. Useful when you want to push commits to some different repo than the one configured in the machine
* Takes 2 arguments: [git_username] and [git_email]

####
```bash
addLocalGitUser sayden mariocaster@gmail.com
```
Will produce `git config --local --add user.name sayden` and `git config --local --add user.email mariocaster@gmail.com`. You can check results running `git config --local -l`
