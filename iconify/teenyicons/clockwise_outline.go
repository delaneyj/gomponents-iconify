package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockwiseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14 8.495v-.5h-1v.5h1ZM7.5 2.999H8v-1h-.5v1Zm1-.5l.353.353l.354-.353l-.354-.354l-.353.354ZM13 8.495a5.499 5.499 0 0 1-5.5 5.496v1c3.589 0 6.5-2.909 6.5-6.496h-1ZM7.5 13.99A5.499 5.499 0 0 1 2 8.495H1a6.499 6.499 0 0 0 6.5 6.496v-1ZM2 8.495a5.499 5.499 0 0 1 5.5-5.496v-1A6.499 6.499 0 0 0 1 8.495h1ZM6.147.854l2 1.998l.706-.707l-2-1.999l-.706.708Zm2 1.291l-2 1.999l.706.707l2-1.999l-.706-.707Z"/>`),
		g.Group(children),
	)
}