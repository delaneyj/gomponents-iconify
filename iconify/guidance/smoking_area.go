package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmokingArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M23.5 11v6m-9-16v1A6.5 6.5 0 0 0 21 8.5m2.5.5V5.5H21A3.5 3.5 0 0 1 17.5 2V1M.5 16.5v-5h20v5H.5Z"/>`),
		g.Group(children),
	)
}