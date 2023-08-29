package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListFlat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 9h9v1H2zm0 3h8v1H2zm0-6h12v1H2zm0-3h11v1H2z"/>`),
		g.Group(children),
	)
}