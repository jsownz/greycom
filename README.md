# greycom v1.0
### Greynoise CLI for Community API users
---
**ONLY TESTED ON 64BIT UBUNTU 20.04**

I'm not a go dev, mostly just pieced some code together to get a cli that would work for the community API (apparently, the API key doesn't even matter for these endpoints, so I changed it to a warning instead of an exit).

Only feature available for community is standard IP lookup, so don't expect anything fancy.

Haven't parsed output yet, just outputs the json.

**Install**

Again, not a go dev, I think this is how you're supposed to do it, who knows - also added a binary release but it was only built on 64 bit linux, so feel free to compile yourself (these instructions would be different for Windows, use your flavor of golang to build and install)
```bash
go get github.com/jsownz/greycom
cd $(go env GOPATH)/src/github.com/jsownz/greycom
go install
sudo ln -s $(go env GOPATH)/bin/greycom /usr/local/bin/greycom
```

**Usage**
```bash
greycom -t [target_ip]
```

To save your Community API key, use the -apikey flag:
```bash
greycom -t [target_ip] -apikey [community_api_key]
```