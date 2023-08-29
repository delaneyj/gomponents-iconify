package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertOctagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.103 2h5.794a3 3 0 0 1 2.122.879l4.101 4.1A3 3 0 0 1 22 9.104v5.794a3 3 0 0 1-.879 2.122l-4.1 4.101a3 3 0 0 1-2.123.88H9.103a3 3 0 0 1-2.122-.88l-4.101-4.1A3 3 0 0 1 2 14.897V9.103a3 3 0 0 1 .879-2.122l4.1-4.101A3 3 0 0 1 9.104 2zM12 8v4m0 4h.01"/>`),
		g.Group(children),
	)
}