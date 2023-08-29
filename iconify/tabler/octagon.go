package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Octagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.103 2h5.794a3 3 0 0 1 2.122.879L21.12 6.98A3 3 0 0 1 22 9.103v5.794a3 3 0 0 1-.879 2.122L17.02 21.12a3 3 0 0 1-2.122.879H9.103a3 3 0 0 1-2.122-.879l-4.101-4.1A3 3 0 0 1 2 14.897V9.103a3 3 0 0 1 .879-2.122L6.98 2.88A3 3 0 0 1 9.103 2z"/>`),
		g.Group(children),
	)
}