package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.995 4.917h1.992V0H1l8.026 2.666v12.251L1.321 13l-.215.829l8.89 2.153v-3.085l1.99.018V9.042h-1.99V4.917z"/><path d="m15.904 7l-2.87-2.932v1.987H11v1.916h2.034v2.062L15.904 7zM7 8h.958v.916H7z"/></g>`),
		g.Group(children),
	)
}