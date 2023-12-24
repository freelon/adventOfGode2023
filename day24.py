lines = open('input/day24.txt').read().splitlines()

arr = []
for line in lines:
    left, right = line.split(' @ ')
    px, py, pz = list(map(int, left.split(', ')))
    vx, vy, vz = list(map(int, right.split(', ')))
    arr.append((px, py, pz, vx, vy, vz))

from math import inf

V = set()

def sol(x1, y1, z1, v1x, v1y, v1z, x2, y2, z2, v2x, v2y, v2z):
    if v1y*v2x-v1x*v2y == 0: return inf, inf
    x = (v1x*v2x*(y2-y1)-v2y*v1x*x2+v1y*v2x*x1)/(v1y*v2x-v1x*v2y)
    y = (v1y*v2y*(x2-x1)-v2x*v1y*y2+v1x*v2y*y1)/(v1x*v2y-v1y*v2x)
    if v1x and (x-x1)/v1x < 0 or v2x and (x-x2)/v2x < 0: return inf, inf
    if v1y and (y-y1)/v1y < 0 or v2y and (y-y2)/v2y < 0: return inf, inf
    return x, y
low = 200000000000000
up = 400000000000000
res = 0
c = 0
for i in range(len(arr)):
    for j in range(i+1, len(arr)):
        c += 1
        x, y = sol(*arr[i], *arr[j])
        print(*arr[i], *arr[j])
        print(x, y)
        if low <= x <= up and low <= y <= up:
#             print(x, y)
#             print(*arr[i], ' ', *arr[j])
            res += 1
            print(i, j, True)
        else:
            print(i, j, False)
print(c)
print(res)