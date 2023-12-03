# Advent of Code in Go

A handy template repository to hold your [Advent of Code](https://adventofcode.com) solutions in Go (golang).

Advent of Code (https://adventofcode.com) is a yearly series of programming questions based on the [Advent Calendar](https://en.wikipedia.org/wiki/Advent_calendar). For each day leading up to christmas, there is one question released, and from the second it is released, there is a timer running and a leaderboard showing who solved it first.

---

### Features

* A directory per question `<year>/<day>`
* Auto-download questions into `<year>/<day>/README.md`
* Auto-download example input into `<year>/<day>/input-example.txt`
* With env variable `AOC_SESSION` set:
   * Auto-download part 2 of questions into `<year>/<day>/README.md`
   * Auto-download user input into `<year>/<day>/input-user.md`
   * Only runs part 2 once part 1 is completed 
* When you save `code.go`, it will execute your `run` function 4 times:
   * Input `input-example.txt` and `part2=false`
   * Input `input-example(2).txt` and `part2=true`
   * Input `input-user.txt` and `part2=false`
   * Input `input-user(2).txt` and `part2=true`
   * Each run will display the return value and timing.
   * Part 2 will use the `<file>2.txt` if it exists.
* Control execution with `PART= INPUT= ./run.sh <year> <day>`, where
   * `PART` can be `1` or `2`, and
   * `INPUT` can be `example` or `user`

---

### Usage

1. Click "**Use this template**" above to fork it into your account
1. Setup repo, either locally or in codespaces
   * Locally
      * Install Go from https://go.dev/dl/ or from brew, etc
      * Git clone your fork
      * Open in VS Code, and install the Go extension
   * Codespaces
      * Click "Open in Codespaces"
1. Open a terminal and `./run.sh <year> <day>` like this:

   ```sh
   $ ./run.sh 2023 1
   Created directory ./2023/01
   Created file code.go
   Created file README.md
   Created file input-example.txt
   run(part1, input-example) returned in 616µs => 42
   ```

1. Implement your solution in `./2023/01/code.go` inside the `run` function
   * I have provided solutions for year `2022`, days `2`,`4`,`7` – however you can delete them and do them yourself if you'd like
1. Changes will re-run the code
   * For example, update `code.go` to `return 43` instead you should see:

   ```sh
   file changed code.go
   run(part1, input-example) returned in 34µs => 43
   ```

1. The question is downloaded to `./2023/01/README.md`
1. Login to https://adventofcode.com
1. Find your question (e.g. https://adventofcode.com/2023/day/1) and **[get your puzzle input](https://adventofcode.com/2023/day/1/input)** and save it to `./2023/01/input-user.txt`
   * See **Session** below to automate this step 
1. Iterate on `code.go` until you get the answer
1. Submit it to https://adventofcode.com/2023/day/1

---

#### Session

**Optionally**, you can `export AOC_SESSION=<session>` from your adventofcode.com `session` cookie. That is:

* Login with your browser
* Open developer tools > Application/Storage > Cookies
* Retrieve the contents of `session`
* Export it as `AOC_SESSION`

With your session set, running `code.go` will download your user-specifc `input-user.txt` and also update `README.md` with part 2 of the question once you've completed part 1.

Currently, your session is NOT used to submit your answer. You still need to login to https://adventofcode.com to submit.