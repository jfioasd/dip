import sys

def run_dip(prog, stack):
    ip = 0

    while ip < len(prog):
        if prog[ip] == '0':
            stack.append(0)

        elif prog[ip] == '!':
            print('(', *stack, ')')
        elif prog[ip] == '.':
            print(stack[-1], end = ' ')

        elif prog[ip] == "'":
            stack[-1] += 1
        elif prog[ip] == ';':
            stack.insert(0, stack.pop())

        elif prog[ip] == '(':
            prev = ip
            lvl = 1
            while ip + 1 < len(prog) and lvl > 0:
                ip += 1
                lvl += prog[ip] == '('
                lvl -= prog[ip] == ')'
            
            while (n := stack.pop()) != 0:
                stack.append(n-1)
                stack = run_dip(prog[prev+1:ip], stack)

        ip += 1
    
    return stack

if __name__ == '__main__':
    prog = open(sys.argv[-1]).read()
    x = [int(i) for i in input().split()]

    print('(', *run_dip(prog, x), ')')
