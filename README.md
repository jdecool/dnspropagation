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
$ go run main.go [-config path/to/configuration/file.hcl] [domain]
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

# Configuration sample

```hcl
dns {
  name = "Cloudflare"

  server {
    protocol = "udp"
    primary = "1.1.1.1:53"
    secondary = "1.0.0.1:53"
  }
}

dns {
  name = "Google"

  server {
    protocol = "udp"
    primary = "8.8.8.8:53"
    secondary = "8.8.4.4:53"
  }
}

dns {
  name = "OpenDNS"

  server {
    protocol = "udp"
    primary = "208.67.222.222:53"
    secondary = "208.67.220.220:53"
  }
}
```
