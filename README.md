# ¡¡¡DISCLAIMER!!!

This is my first serious project hope it helps.

This software is **not affiliated with [Castle: Make and Play](https://castle.xyz)** this is an independent and open source library for educational or deck-making purposes.
Im not affiliated with [maksydab](https://castle.xyz/@maksydab) hes a cool person.

**USE AT YOUR OWN RISK!!** Based in ToS the high use of automated tools can be **banned**!! [ToS](https://castle.xyz/terms)



...

It needs an auth token (get it seeing headers in console>net>headers X-Auth-Token) the information is not shared as you can see in the main go (castle-go-net.go)

**REMIND!** If you gonna make tons of requests, please, add a time limit!

Also, dont share ur auth token to anyone!

Im not responsible of any bad action made with this tool.



# How to install

GOPROXY direct: <br>

`Linux: GOPROXY=direct go get -u github.com/clfurry/castlenet/` <br>
`Windows Powershell: $env:GOPROXY="direct"; go get -u github.com/clfurry/castlenet/` <br>
`Windows CMD: set GOPROXY=direct && go get -u github.com/clfurry/castlenet/` <br>
MacOS: Same as linux

Normal: <br>

`go get -u github.com/clfurry/castlenet/`

MADE IN GO 1.22!!!!

# WIKI

`castlenet.LoadAuth(xauth string) > Nothing: Loads the auth token (Necessary to modify and read shared storage`

`castlenet.InternalPostHttp(url string, payload []byte) > response []byte, err error: Utility (forgot to make it internal)`

`castlenet.DecodeJson(body []byte, wheretosave Pointer any) > err error: Utility (Not used)`

`castlenet.GetSharedData(deckId string, toVar string) > got any, err error: toVar is necessary for castle so it knows what it needs to get, deckId too (Get both with console)`

`castlenet.UpdateSharedData(deckId string, toVar string, val string) > got any, err error: same as get but it modify the data, everything even numbers are string in the request`

Getting deckid:
`go to the page wheres your deck and copy example: https://castle.xyz/d/Ohd78EX2IM5y/lua-test > Ohd78EX2IM5y`

Types: <br>
`For numbers: "0"` <br>
`For strings: "\\\"Hello World!\\\""` <br>
`For arrays: "[\\\"hi\\\",0]"` <br>
`For tables: "{\\\"dam\\\":0,\\\"hello\\\":[\\\"h\\\",\\\"e\\\",\\\"l\\\",\\\"l\\\",\\\"o\\\"],\\\"thing\\\":\\\"hi\\\"}"` <br>
`For booleans: "false"/"true"` <br>
`For nil: "null"` <br>
`For mixed table: "{\\\"1\\\":\\\"ha\\\",\\\"2\\\":\\\"he\\\",\\\"3\\\":\\\"hi\\\",\\\"ho\\\":\\\"ho\\\",\\\"hu\\\":\\\"hu\\\"}"`
