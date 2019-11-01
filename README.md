
# What is this ?
This is an Alfred workflow that contains some utils for developers :
* Show IP list of local network interfaces.
* Search location of an input IP.
* MD5 an input string.
* BASE64 encode/decode input string.
* UrlEncode/UrlDecode input string.
* Generate struct from json string for go/GoLang.


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



# Thanks
* https://github.com/upx/upx