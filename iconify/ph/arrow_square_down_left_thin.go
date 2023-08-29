package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareDownLeftThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 36H48a12 12 0 0 0-12 12v160a12 12 0 0 0 12 12h160a12 12 0 0 0 12-12V48a12 12 0 0 0-12-12Zm4 172a4 4 0 0 1-4 4H48a4 4 0 0 1-4-4V48a4 4 0 0 1 4-4h160a4 4 0 0 1 4 4ZM162.83 93.17a4 4 0 0 1 0 5.66L105.66 156H144a4 4 0 0 1 0 8H96a4 4 0 0 1-4-4v-48a4 4 0 0 1 8 0v38.34l57.17-57.17a4 4 0 0 1 5.66 0Z"/>`),
		g.Group(children),
	)
}