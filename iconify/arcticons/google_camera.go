package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.6 9.54v2.11H6.29a1.79 1.79 0 0 0-1.79 1.79h0v23.23a1.79 1.79 0 0 0 1.79 1.79h35.42a1.79 1.79 0 0 0 1.79-1.79h0V13.45a1.8 1.8 0 0 0-1.79-1.8H16.39V9.54ZM24 17.75a8.52 8.52 0 1 1-8.51 8.51A8.51 8.51 0 0 1 24 17.75Z"/>`),
		g.Group(children),
	)
}