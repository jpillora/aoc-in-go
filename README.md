# Advent of Code in Go

A handy template repository to hold your [Advent of Code](https://adventofcode.com) solutions in Go (golang).

Advent of Code (https://adventofcode.com) is a yearly series of programming questions based on the [Advent Calendar](https://en.wikipedia.org/wiki/Advent_calendar). For each day leading up to christmas, there is one question released, and from the second it is released, there is a timer running and a leaderboard showing who solved it first.

### Features

* Structured questions into `<year>/<day>`
* Auto-download questions into `<year>/<day>/README.md`
* Auto-download example input into `<year>/<day>/input-example.txt`
* With env variable `AOC_SESSION` set:
   * Auto-download part 2 of questions into `<year>/<day>/README.md`
   * Auto-download user input into `<year>/<day>/input-user.md`
* When you save `code.go`, it will execute your `run` function 4 times:
   * Input `input-example.txt` and `part2=false`
   * Input `input-example.txt` and `part2=true`
   * Input `input-user.txt` and `part2=false`
   * Input `input-user.txt` and `part2=true`
   * and, will show the results and timing of each

### Usage

* Click "Use this template" above to fork it into your account
* Setup repo, either locally or in codespaces
   * Locally
      * Install Go from https://go.dev/dl/ or from brew, etc
      * Git clone your fork
      * Open in VS Code, and install the Go extension
   * Codespaces
      * Click "Open in Codespaces"
* Open a terminal and `./run.sh <year> <day>`:

   ```sh
   ./run.sh 2022 1
   [run.sh] created ./2022/01
   [run.sh] created ./2022/01/code.go
   Created file README.md
   Created file input-example.txt
   run(part1, input-example) returned in 616µs => 42
   ```

* Implement your solution in `./2022/01/code.go` inside the `run` function
   * I have provided solutions for year `2022`, days `2`,`4`,`7` – however you can delete them and do them yourself if you'd like
* Changes will re-run the code
   * For example, `update` `return 43` instead you should see

   ```sh
   file changed code.go
   run(part1, input-example) returned in 34µs => 43
   ```

* The question is downloadded to `README.md`, iterate on `code.go` until you get the answer
* Login to https://adventofcode.com and submit


#### Session

**Optionally**, you can set `export AOC_SESSION=<session>` to your adventofcode.com `session` cookie. That is:

* Login with your browser
* Open developer tools > Application/Storage > Cookies
* Retrieve the contents of `session`
* Export it as `AOC_SESSION`

With your session, `puzzler` will download your user-specifc `input-user.txt` and also update `README.md` with part 2 of the question once you've completed part 1.

Current, your session is NOT used to submit your answer. You still need to login to https://adventofcode.com to submit.