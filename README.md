# tmall-rate-crawler

天猫商品评论数据拉取工具

## Introduction

Tmall products rating data crawler. （being developed）

### Features

- Fetch rating data

### TODO

- Implement data analyzer

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
   rate-crawler fetch - fetch specified item's rate data

USAGE:
   rate-crawler fetch [command options] [arguments...]

OPTIONS:
   --user-agent value        The HTTP request user-agent header (default: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")
   --cookie value, -c value  The cookie data
   --item value, -i value    The item ID
   --seller value, -s value  The seller ID
   --interval value          The request interval(Second) (default: 3)
   --start value             The start page number (default: 1)
   --output value, -o value  The output file name (default: "result.csv")
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