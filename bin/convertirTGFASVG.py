from graphviz import Digraph
import os
import sys

print sys.argv


dot = Digraph(comment='Archivo TGF convertido a SVG')
dot

lineas = [lineas.rstrip('\n') for lineas in open(os.getcwd()+'/'+sys.argv[1], 'r')]

iniciaEdge = False

for linea in lineas:
	contenido = linea.split()
	if contenido[0] == '#':
		iniciaEdge = True
		continue
	if iniciaEdge:
		dot.edge(contenido[0], contenido[1], constraint='true')
	else:
		dot.node(contenido[0], contenido[1])

dot.format = 'svg'
dot.render('output.gv', view=True)
print lineas