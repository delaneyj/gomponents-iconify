package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListNumberedMirror(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 28h-6v-4c0-1.1.9-2 2-2h2v-2h-4v-2h4c1.1 0 2 .9 2 2v2c0 1.1-.9 2-2 2h-2v2h4v2zM2 22h14v2H2zm24-10V4h-2v1h-2v2h2v5h-2v2h6v-2zM2 8h14v2H2z"/>`),
		g.Group(children),
	)
}