package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Contacter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M3 5H2V4h1v1zm13.038-1.958H2.046c-.556 0-1.004.452-1.004 1.012v8.976c0 .56.448 1.012 1.004 1.012h13.992c.556 0 1.004-.451 1.004-1.012V4c0-.56-.448-.958-1.004-.958zM3 13H2v-1h1v1zm11-2H4V6h10v5zm2 2h-1v-1h1v1zm0-8h-1V4h1v1z"/><path d="M5 7v3h5V7H5z"/></g>`),
		g.Group(children),
	)
}