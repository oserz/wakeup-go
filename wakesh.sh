#!/bin/sh
MAC_ADDR="01-02-03-04-05-06"
WAKEUP_BIN="wakeup"
LOCALIP_ADDR=$(ifconfig | awk 'NR==1 {if (index($1,":") == 0) print $1; else print substr($1, 0, index($1,":")-1)}')
Basepath=$(cd `dirname $0`; pwd)
$Basepath"/"$WAKEUP_BIN -mac "$MAC_ADDR" -interface "$LOCALIP_ADDR"