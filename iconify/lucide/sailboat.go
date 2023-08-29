package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sailboat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 18H2a4 4 0 0 0 4 4h12a4 4 0 0 0 4-4Zm-1-4L10 2L3 14h18ZM10 2v16"/>`),
		g.Group(children),
	)
}