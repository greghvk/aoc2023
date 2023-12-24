import sympy

x, y, z, vx, vy, vz = sympy.symbols("x, y, z, vx, vy, vz")

f = open("./input", 'r')
lines = [l.replace(" @ ", ", ").split(", ") for l in f.readlines()]

eqs = []
for line in lines[:5]:
  [xi, yi, zi, vxi, vyi, vzi] = [int(i) for i in line]
  # Equation derivation:
  # x + vx*t = xi + vxi*t
  # x - xi = t * (vxi - vx)
  # t = (x-xi) / (vxi - vx) # repeat the same for y and z, you end up with 3 equations that are equal:
  # 1st (x-xi) / (vxi - vx)
  # 2nd (y-yi) / (vyi - vy)
  # 3rd (z-zi) / (vzi - vz)

  # Created by equating 1st and 2nd (moving everything to one side of "="):
  eqs.append((x-xi)*(vyi - vy) - (y-yi)*(vxi-vx))

  # And by equating 2nd and 3rd:
  eqs.append((y-yi)*(vzi - vz) - (z-zi)*(vyi-vy))

res = sympy.solve(eqs)
print(res[0][x] + res[0][y] + res[0][z])


  


