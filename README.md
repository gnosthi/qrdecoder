# QRDecoder 

QRDecoder is, surprise, surprise; A QR Image Decoder written in go.

## Why was this app written? 
With a cursory glance of a google search, I couldn't find a cli QR image decoder. I was trying to decode a QR image for an OTP service
and found that I wasn't able to find a quick and easy QR image decoder. All the decoders I was able to find were web services. I wanted something
that was quick, simple, and relatively secure. Passing an OTP QR code into a webservice is not secure. Most of the QR code scanner sites
are not simple, and none of the tools worked on a command line.

As a result I wrote this up in about 15 minutes, thanks to https://github.com/liyue201/goqr. 

Of course this tool can be expanded a great deal, but for the time being it met my purposes exactly. 

## Usage

Simply call the application with a file, or list of files you want it to operate on. 

```bash
qrdecode file1 file2 file3 ... 
```
The output will look something like this. 

```bash
./qrdecode ./qrcode.png ./qrcode2.png
File: /home/account/location/qrcode.png
 Hello World! 
File: /home/account/location/qrcode2.png
 Hello Again! 
```
