dnspropagation
==============

A simple tool to check DNS state of a domain.

## Usage

Install it:

```bash
$ composer install
```

Then use it:

```bash
$ php bin/propagation check [domain]
```

You can also specify the exepected DNS resolution by using the `--expected` option:

```bash
$ php bin/propagation check [domain] --expected=[ip]
```
