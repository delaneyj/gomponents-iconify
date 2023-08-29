package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogicNand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12h-3M2 9h3m-3 6h3M7 5c6 0 8 3.5 8 7s-2 7-8 7H5V5h2zm8 7a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/>`),
		g.Group(children),
	)
}