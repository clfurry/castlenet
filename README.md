# ¡¡¡DISCLAIMER!!!

This software is **not affiliated with [Castle: Make and Play](https://castle.xyz)** this is an independent and open source library for educational purposes.

**USE AT YOUR OWN RISK!!** Based in ToS the high use of automated tools can be **banned**!! [ToS](https://castle.xyz/terms)



...

It needs an auth token (get it seeing headers in console>net>headers X-Auth-Token) the information is not shared as you can see in the main go (castle-go-net.go)

**REMIND!** If you gonna make tons of requests, please, add a time limit!

Also, dont share ur auth token to anyone!

Im not responsible of any bad action made with this tool.

# WIKI

`castlenet.LoadAuth(xauth string) > Nothing: Loads the auth token (Necessary to modify and read shared storage`

`castlenet.InternalPostHttp(url string, payload []byte) > response []byte, err error: Utility (forgot to make it internal)`

`castlenet.DecodeJson(body []byte, wheretosave Pointer any) > err error: Utility (Not used)`

`castlenet.GetSharedData(deckId string, toVar string) > got any, err error: toVar is necessary for castle so it knows what it needs to get, deckId too (Get both with console)`

`castlenet.UpdateSharedData(deckId string, toVar string, val string) > got any, err error: same as get but it modify the data, everything even numbers are string in the request`

Getting deckid:
`go to the page wheres your deck and copy example: https://castle.xyz/d/Ohd78EX2IM5y/lua-test > Ohd78EX2IM5y`

Types:
`For numbers: "0"`
`For strings: "\\\"Hello World!\\\""`
`For tables: [\\\"hi\\\",0]`
