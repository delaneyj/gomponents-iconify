package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitlabOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m.5 8.5l7 6l7-6l-2-8l-2 6h-6l-2-6l-2 8Z"/>`),
		g.Group(children),
	)
}