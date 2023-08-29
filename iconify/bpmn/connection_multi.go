package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectionMulti(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1876.66 106.365l-840.543 224.07l130.275 193.138l-995.052 671.175l51.399 76.183l995.051-671.175l126.692 187.846l532.162-681.227zM250.133 1812.293l-78.793 53.156l51.398 76.193l78.792-53.147zm1092.706-352.887L1875 778.188l.016-.01l-840.542 224.06zm-304.962-178.451l-128.523 86.69l51.398 76.193l128.523-86.68zm-257.055 173.381l-128.53 86.69l51.395 76.213l128.533-86.7zm-273.634 184.576l-128.532 86.7l51.397 76.202l128.533-86.72z"/>`),
		g.Group(children),
	)
}