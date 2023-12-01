#!/bin/bash
set -euf -o pipefail

# functions
function template() {
    cat <<EOF
package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	if part2 {
		return "not implemented"
	}
	return 42
}
EOF
}

# two args YEAR and DAY
YEAR="${1:-}"
DAY="${2:-}"
if [ -z "$YEAR" ] || [ -z "$DAY" ]; then
    echo "Usage: $0 <YEAR> <DAY>"
    exit 1
fi
# pad DAY to 2 digits
DAY=$(printf "%02d" $DAY)
DIR="./$YEAR/$DAY"
# create missing files as needed
if [ ! -d "$DIR" ]; then
    mkdir -p "$DIR"
    echo "[run.sh] created $DIR"
fi
if [ ! -f "$DIR/code.go" ]; then
    template >"$DIR/code.go"
    echo "[run.sh] created $DIR/code.go"
fi
# go run
cd "$DIR" && go run code.go
