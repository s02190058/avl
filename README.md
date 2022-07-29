# Generic performance in Go

I was interested to see how to implement generalized types in Go 1.18.
I also wanted to compare their performance with the implementation for specific types.
So I wrote a simple AVl tree (generic, `int` and `string` version)
and a benchmark that adds elements to an empty tree and then cleans it completely.

## Results

**int tree:**

| size  | certain implementation performance (ns/op) | generic performance (ns/op) |
|:-----:|:------------------------------------------:|:---------------------------:|
|  32   |                    2750                    |            2936             |
| 1024  |                   247502                   |           241176            |
| 32768 |                  15153595                  |          13542495           |

**string tree:**

| size  | certain implementation performance (ns/op) | generic performance (ns/op) |
|:-----:|:------------------------------------------:|:---------------------------:|
|  32   |                    4279                    |            4032             |
| 1024  |                   366448                   |           337257            |
| 32768 |                  23586607                  |          20246445           |

## Summary

Implementation through generalized types does not lose at all to implementations with specific types. 
On my computer, generics work faster on almost all amounts of data.
Using square brackets to specify type parameters is a bit unusual, but also convenient.
