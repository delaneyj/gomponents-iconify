package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowURightDownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m221.66 181.66l-48 48a8 8 0 0 1-11.32 0l-48-48A8 8 0 0 1 120 168h40V88a48 48 0 0 0-96 0v88a8 8 0 0 1-16 0V88a64 64 0 0 1 128 0v80h40a8 8 0 0 1 5.66 13.66Z"/>`),
		g.Group(children),
	)
}