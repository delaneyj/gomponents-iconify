package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListSelection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1 12v-1h9v1H1zm0-5h14v1H1V7zm11-4v1H1V3h11z"/>`),
		g.Group(children),
	)
}