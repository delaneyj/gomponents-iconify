package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpVideoFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 2H4v20h16V8l-6-6zm-1 7V3.5L18.5 9H13zm1 5l2-1.06v4.12L14 16v2H8v-6h6v2z"/>`),
		g.Group(children),
	)
}