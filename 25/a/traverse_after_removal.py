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

S = [G.subgraph(c).copy() for c in nx.connected_components(G)]
print(S[0].order() * S[1].order())

