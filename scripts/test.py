import sys


print('第一行stdout')
print('第二行stderr', file=sys.stderr)
print('第三行stdout')
print('第四行stderr', file=sys.stderr)
a = 1 + 'b'