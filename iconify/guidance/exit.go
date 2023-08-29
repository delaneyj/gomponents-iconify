package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M24 12H8m10.5-5c0 .577.665 1.562 1.228 2.294a7.494 7.494 0 0 0 1.745 1.662C22.2 11.445 23.2 12 23.99 12c-.79 0-1.79.556-2.517 1.044a7.494 7.494 0 0 0-1.745 1.662c-.563.732-1.228 1.717-1.228 2.294m-4-10V2.5h-.329A46 46 0 0 1 1.103.605L.75.5H.5v23h.25l.353-.105A45.998 45.998 0 0 1 14.171 21.5h.329V17"/>`),
		g.Group(children),
	)
}