# go-wyag
## Write Yourself a Git in Go

I created this project to learn git internals and golang better after I read and got inspired by [this documentation](https://wyag.thb.lt/#org4a4112c).  

It is basically a Git client written from scratch and implements some of Git's functionality like log, commit, hash-object etc. I would definitely recommend anyone who is interested to understand Git better to read the linked documentation and write your own Git! Not only did I learn Git better, but I was also challenged to think about how to structure the code, what abstractions to use to minimize duplication etc.

## wyag log command

The log command, the equivalent of git log, is for simplicity sake rendered into a [dot graph](https://en.wikipedia.org/wiki/DOT_(graph_description_language)) as a pdf to show the dependencies between the commits. The following example uses graphviz to create these files.

You can test it with:
```
make -s log > log.dot
dot -O -Tpdf log.dot
open log.dot.pdf
```