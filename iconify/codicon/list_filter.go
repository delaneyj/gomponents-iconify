package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListFilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 12v-1h4v1H6zM4 7h8v1H4V7zm10-4v1H2V3h12z"/>`),
		g.Group(children),
	)
}