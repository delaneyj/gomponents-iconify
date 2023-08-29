package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarricadeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 68H32a12 12 0 0 0-12 12v72a12 12 0 0 0 12 12h28v36a4 4 0 0 0 8 0v-36h120v36a4 4 0 0 0 8 0v-36h28a12 12 0 0 0 12-12V80a12 12 0 0 0-12-12Zm4 12v58.34L165.66 76H224a4 4 0 0 1 4 4ZM32 76h50.34l80 80h-60.68L28 82.34V80a4 4 0 0 1 4-4Zm-4 76V93.66L90.34 156H32a4 4 0 0 1-4-4Zm196 4h-50.34l-80-80h60.68L228 149.66V152a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}