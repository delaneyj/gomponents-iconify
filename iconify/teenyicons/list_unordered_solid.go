package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListUnorderedSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 4H0V3h2v1Zm13 0H4V3h11v1ZM2 8H0V7h2v1Zm13 0H4V7h11v1ZM2 12H0v-1h2v1Zm13 0H4v-1h11v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}