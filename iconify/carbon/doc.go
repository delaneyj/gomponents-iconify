package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Doc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 23h-6a2.002 2.002 0 0 1-2-2V11a2.002 2.002 0 0 1 2-2h6v2h-6v10h6zm-12 0h-4a2.002 2.002 0 0 1-2-2V11a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v10a2.002 2.002 0 0 1-2 2zm-4-12v10h4V11zM6 23H2V9h4a4.005 4.005 0 0 1 4 4v6a4.005 4.005 0 0 1-4 4zm-2-2h2a2.002 2.002 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2H4z"/>`),
		g.Group(children),
	)
}