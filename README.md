dnspropagation
==============

A simple tool to check DNS state of a domain.

This tool as completly been rewrite from [PHP](https://github.com/jdecool/dnspropagation/tree/php).

## Usage

Compile it:

```bash
$ go build
```

Then use it:

```bash
$ ./dnspropagation [domain]
```

You can also run it without compilation:

```
$ go run main.go [domain]
```

# Output

```text
Provider          Primary         Secondary
--------          -------         ---------
Bouygues Telecom  140.82.121.4    140.82.121.4
Cloudflare        140.82.121.3    140.82.121.4
Comodo            20.205.243.166  20.205.243.166
FDN               140.82.121.4    140.82.121.3
FreeDNS           140.82.121.4    140.82.121.4
Google            140.82.121.4    140.82.121.3
Neustar           140.82.121.3    140.82.121.3
Norton            140.82.121.3    140.82.113.4
OpenDNS           140.82.121.3    140.82.121.3
Quad9             140.82.112.3    140.82.112.4
Verisign          140.82.113.4    140.82.121.4

Time: 5.03 seconds
```
