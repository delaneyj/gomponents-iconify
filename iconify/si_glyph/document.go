package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Document(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M2 0v16h13V4.024L9 4V.062L2 0zm11 13H4v-1h9v1zm0-2H4v-1h9v1zm0-4v1H4V7h9z"/><path d="M10 0v2.844h4.752L10 0z"/></g>`),
		g.Group(children),
	)
}