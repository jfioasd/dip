# dip
A tiny esolang

## Commands
| Command | Behavior |
| :-:     | :-:      |
| `0`     | Push `0`. |
| `'`     | Increment TOS. |
| `;`     | Pop TOS. Append it to the bottom of the stack. |
| `( _b_ )` | While loop. (see below) |
### While loop:
* While (pop stack as `N`) != `0`:
  * push `N-1`
  * run `_b_` as dip code

## Extra commands & I/O
For I/O, the program reads in a line of space-separated integers. This is used as the initial stack.

At the end of the program, the whole stack is outputted (like there's a `!` command).

Here are the extra commands in the interpreter:
* `!`: Print the entire stack.
* `.`: Print TOS as a number. (Does not pop the stack.)
