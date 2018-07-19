#!/bin/bash
MYPWD=${PWD}

if [ "$1" = "start" ]; then

        echo -n "Starting proxy on http://127.0.0.1:$2 ..."
        /bin/sh -c " $MYPWD/bin/go-proxy $2 " &> log/$2.log &
        echo "Done!"

elif [ "$1" = "stop" ]; then

        echo -n "Stoping proxy on http://127.0.0.1:$2 ..."
        kill -9 $(lsof -t -i:$2 -sTCP:LISTEN)
        echo "Done!"
else
        echo "Error: usage {start|stop} + {port}"
fi