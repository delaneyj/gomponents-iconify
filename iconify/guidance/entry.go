package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Entry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M17 12H1m10.5-5c0 .577.665 1.562 1.228 2.294a7.494 7.494 0 0 0 1.745 1.662C15.2 11.445 16.2 12 16.99 12c-.79 0-1.79.556-2.517 1.044a7.494 7.494 0 0 0-1.745 1.662c-.563.732-1.228 1.717-1.228 2.294m-3-10V2.5h.329A46 46 0 0 0 21.897.605L22.25.5h.25v23h-.25l-.353-.105A45.998 45.998 0 0 0 8.829 21.5H8.5V17"/>`),
		g.Group(children),
	)
}