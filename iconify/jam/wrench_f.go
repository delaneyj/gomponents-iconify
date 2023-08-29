package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrenchF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.477 13.84l-4.243 4.243L8.01 11.86a6.002 6.002 0 0 1-6.445-9.273L5.1 6.12a1 1 0 0 0 1.415-1.414L2.979 1.17a6.002 6.002 0 0 1 9.273 6.445l6.225 6.225zm1.414 1.415l.707.707a3 3 0 0 1-4.243 4.243l-.707-.708l4.243-4.242z"/>`),
		g.Group(children),
	)
}