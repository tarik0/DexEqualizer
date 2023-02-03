#!/bin/bash

var="$(pm2 list | grep dexeq_out | grep "stopped" | wc -c)"

if [[ $var =~ "0" ]]; then
    pm2 stop dexeq_out && pm2 start dexeq_out
fi