package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileExcel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3V1Zm2 2v18h14V9h-6V3H5Zm10 .414V7h3.586L15 3.414ZM15 10v1.646a2 2 0 0 1-.612 1.44l-.947.914l.947.914a2 2 0 0 1 .612 1.44V18h-2v-1.646l-1-.965l-1 .965V18H9v-1.646a2 2 0 0 1 .612-1.44l.947-.914l-.947-.914A2 2 0 0 1 9 11.646V10h2v1.646l1 .965l1-.965V10h2Z"/>`),
		g.Group(children),
	)
}