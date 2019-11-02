
# What is this ?
This is an Alfred workflow that contains some utils for developers :
* Show IP list of local network interfaces.
* Search location of an input IP.
* MD5 an input string.
* BASE64 encode/decode input string.
* UrlEncode/UrlDecode input string.
* Generate struct from json string for go/GoLang.
* Generate json string from K:V string 

```json
--------- input --->
a:b
c:1
d:
h:1
l:1.1
m:1.1
o:
--------- generated --->
{
  "a": "b",
  "c": "1",
  "d": "",
  "h": "1",
  "l": "1.1",
  "m": "1.1",
  "o": ""
}
```
* Generate K:V string from json string 

```json
--------- input --->
{
  "a": "b",
  "c": "1",
  "d": "",
  "h": "1",
  "l": "1.1",
  "m": "1.1",
  "o" :{
    "p" : "1",
    "q": {
      "x" : "y"
    }
  }
}

--------- generated --->
a:b
c:1
d:
h:1
l:1.1
m:1.1
o:
```


# Install
[download](https://raw.githubusercontent.com/yuanmomo/alfred-utils-go/master/super-momo-tools.alfredworkflow) and double click to install

# Usage
## ip && lip
![ip](https://img.tupm.net/2019/10/F971793758D0D96E9CEC6BBA8CF942B4.jpg)
![lip](https://img.tupm.net/2019/10/907D85B8C8D67D7EDCF9EDA67067C3A8.jpg)

## md
![md5](https://img.tupm.net/2019/10/C8829AB4B8C512EF24BE5F08455CC906.jpg)

## en && de
![base64-encode](https://img.tupm.net/2019/10/A63385AFB1130D9CBB2F623161221ECB.jpg)
![base64-decode](https://img.tupm.net/2019/10/FB07D3CE3CA62F9AAD81E80C8E74B233.jpg)

## uen && ude
**space will encode to %20**

![url-encode](https://img.tupm.net/2019/10/A32F6529D53FE84387AFE81812FB7723.jpg)
![url-decode](https://img.tupm.net/2019/10/0D46DCE2E4CF260074AAE54EC328C2F3.jpg)


## gs && cgs
cgs command will paste json string from clipboard for generating.

### gs
1. copy json 
2. call Alfred window, input gs command, then input json string(press command + v)
3. press enter to copy generated struct to clipboard.

![gs-json-to-struct](https://img.tupm.net/2019/11/61FAC0EA3DFFCEA73A85E4EABA8F4E51.jpg)

### cgs
1. copy json 
2. call Alfred window, input cgs command 
3. press enter to copy generated struct to clipboard.

![cgs-json-to-struct](https://img.tupm.net/2019/11/963F1A33D3211CDD11F4E65370A5A3B0.jpg)



## js && cjs
cjs command will paste K:V string from clipboard for generating.

### js
1. copy K:V string
2. call Alfred window, input js command, then input K:V string(press command + v)
3. press enter to copy generated json string to clipboard.

![kv-to-json](https://img.tupm.net/2019/11/02EA55CB7AB38816049D23923B23F047.jpg)

### cjs
1. copy K:V string 
2. call Alfred window, input cjs command 
3. press enter to copy generated json string to clipboard.

![clipboard-kv-to-json](https://img.tupm.net/2019/11/8F0316063912A64EBC2F2599F79A5BAE.jpg)

## kv && ckv
ckv command will paste K:V string from clipboard for generating.

### kv
1. copy json string
2. call Alfred window, input kv command, then input json string(press command + v)
3. press enter to copy generated K:V string to clipboard.

![json-to-kv](https://img.tupm.net/2019/11/AC9134A66802A3581EE11A8B90DD8370.jpg)


### ckv
1. copy json string 
2. call Alfred window, input ckv command 
3. press enter to copy generated K:V string to clipboard.

![clipboard-json-to-kv](https://img.tupm.net/2019/11/197F5C0338B35E69D3EBE83F299AC70E.jpg)


# Thanks
* https://github.com/upx/upx