# jmattheis/website

This may seems like just a normal website, but it is much more. 
This server abuses various protocols that they will transmit data of my website.

Currently supported are: http/https, websocket, telnet/tcp, whois, dns(tcp), ftp and ssh

Try one of the following commands:


```
curl http://jmattheis.de
curl  ftp://jmattheis.de
dig        @jmattheis.de +tcp +short
netcat      jmattheis.de 23
ssh         jmattheis.de
telnet      jmattheis.de
whois -h    jmattheis.de .
wscat -c    jmattheis.de
```