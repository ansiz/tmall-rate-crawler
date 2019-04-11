# tmall-rate-crawler

天猫商品评论数据拉取工具

## Introduction

Tmall products rating data crawler. （being developed）

### Features

- Fetch comment data
- Fetch shop items data

### TODO

- Automatically fetch comment data based on product id
- Implement data analyzer
- Chrome extension plugin

## Quick Start

### Prerequisite

- Golang 1.8+ installed
- Glide installed

### Installation

```bash
git clone https://github.com/ansiz/tmall-rate-crawler
cd tmall-rate-crawler
glide install
make install
```

### Usage

```txt
NAME:
   Tmall rate crawler - Crawl rate data from Tmall

USAGE:
   rate-crawler [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     rate     fetch specified item's rate data
     item     fetch shop items data
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --verbose                 run in verbose mode
   --user-agent value        The HTTP request user-agent header (default: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")
   --cookie value, -c value  The cookie data
   --interval-min value      The minion request interval(Second) (default: 5)
   --interval-max value      The max request interval(Second) (default: 15)
   --help, -h                show help
   --version, -v             print the version
```

### Example

```bash
rate-crawler --verbose fetch -c 'set your cookie data here' -i 545414402527 -s 2635590370

DEBU[0000] request URI:https://rate.tmall.com/list_detail_rate.htm?currentPage=1&itemId=545414402527&sellerId=2635590370
DEBU[0000] Task(item:545414402527, seller:2635590370) progress: 1/15
DEBU[0003] request URI:https://rate.tmall.com/list_detail_rate.htm?currentPage=2&itemId=545414402527&sellerId=2635590370
DEBU[0003] Task(item:545414402527, seller:2635590370) progress: 2/15
...
```