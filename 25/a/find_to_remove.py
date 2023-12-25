import networkx as nx
import time
import matplotlib.pyplot as plt
import math


G = nx.Graph()

f = open("./25/a/input", 'r')
for l in f.readlines():
  els = l.split(": ")
  src = els[0]
  tgts = els[1]
  for t in tgts.split():
    G.add_edge(src, t)

pos = nx.spring_layout(G)
nx.draw(G, pos, with_labels=True)
plt.show()

