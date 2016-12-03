from graphviz import Digraph
import os
import sys

print sys.argv


dot = Digraph(comment='Archivo TGF convertido a SVG')
dot

iniciaEdge = False

for linea in sys.stdin:
	contenido = linea.split()
	if contenido[0] == '#':
		iniciaEdge = True
		continue
	if iniciaEdge:
		dot.edge(contenido[0], contenido[1], constraint='true')
	else:
		dot.node(contenido[0], contenido[1])

dot.format = 'svg'
print(dot.pipe().decode('utf-8'))