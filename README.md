# dailyhelper
A set of functions used in the daily life of a go programmer.

# How to user
## installation
get the library
```
go get github.com/shouva/dailyhelper
```
## using
```
// import
import helper "github.com/shouva/dailyhelper"
...

//getcurrent dir
currentdir := helper.GetCurrentPath()

//get config

//create struct
struct Setting {
    Name string `json:"name"`
    Nomor uint16 `json:"nomor"`
}

//read the config
var setting Setting
err :=helper.ReadConfig(currentdir + "/config.json", &setting)

fmt.Println(setting.Name)

```

# Sponsorship

This repository is sponsored by [Shouva Store](https://shouva.com) and [Otoritech](https://otoritech.com).