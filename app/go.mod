module github.com/JamesUmmec/CaesarCipher/app

go 1.16

require github.com/JamesUmmec/CaesarCipher/local_server v0.0.0
replace github.com/JamesUmmec/CaesarCipher/local_server => ../server

// this one is the package itself,
// in order to call package in relative path out of GOPATH
require github.com/JamesUmmec/CaesarCipher/main v0.0.0
replace github.com/JamesUmmec/CaesarCipher/main => ./
