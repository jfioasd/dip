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
