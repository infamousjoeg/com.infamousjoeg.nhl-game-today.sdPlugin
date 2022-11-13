#!/bin/sh

pid=$(pgrep -f '/Applications/Stream Deck.app/Contents/MacOS/Stream Deck')
if [ -n "$pid" ]; then
    kill "$pid"
else
    echo "Does not exist"
fi