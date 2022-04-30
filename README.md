<h1 align="center">TURNIEJ TRÓJGAMICZNY</h1>
<a href="https://t3g.pl">
        <p align="center">
                <img src="https://www.t3g.pl/wp-content/uploads/cropped-TTduzy-1.png">
        </p>
</a>

# DESCRIPTION

This module contains a simple, single-file golang app that was
used to conduct a survey about [S(CHEM)E game](http://scheme.noip.pl)
that was neccessary in one of [Tórniej Trójgamiczny](https://t3g.pl)'s tasks.

Results could be viewed in [#1](https://github.com/neonKnights/sCHEMe-poll/issues/1)

The parser package is just a parser to make converting data from  #1 to CSV file easier.

You'll not probably want to use it however, feel free to copy code from here :grinning:

# COMPILATION INSTRUCTION

- install [GO](https://golang.org)
- download the repository `git clone git@github.com:neonKnights/sCHEMe`
- "cd" it `cd sCHEMe-poll`
- edit main.go in order to change default comments destination _xd_
- create file named `token.txt` and post your GitHub Personal Access Token there
- run `go build .`
- You're done!

to parse your data:
- open issue with poll results in your browser
- hop into sCHEMe-poll/parser
- run go app: `go build .`
- copy-paste each comment into your shell, separate them with "ENTER"
- type `$` after the last one
- you're done! output.csv contains parsed data.
  now you can use MS Excel or LibreOffice Calc to open your data
  (optionally save them to something else like ods or xls format)

Is it good solution for the task "Conduct a survy"? Of course it isn't perfect.
Does it look better than Google or MS Forms? No too. It maybe looks
a bit more proffesional like "Run this .exe file and fill form."
but... we're programmers, we want to play with code to learn something.
For example this repo gave me an idea: maybe some day, after end of T3G
I'll write some framework for such a surveys...

# CONTRIBUTION

As I wrote, I'm not going to develop this project as it was just a single-task
project that succeed. I don't recommend to re-use this anyway.
However if you felt some need of contribution, you would be welcome!
