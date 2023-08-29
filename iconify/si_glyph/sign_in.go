package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.959 4.917V1H1.096L9 3.666v12.251L1.219 14l-.215.829L9.959 17v-3.085h2v-2.873l-2.865-3l2.865-3.125z"/><path d="m11.1 8.102l2.87 2.931V8.968h2.046V6.98H13.97V5.068L11.1 8.102zM7 9h.958v.916H7z"/></g>`),
		g.Group(children),
	)
}