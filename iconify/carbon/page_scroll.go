package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageScroll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M16 28H4a1.89 1.89 0 0 1-2-2V14a1.89 1.89 0 0 1 2-2h12a1.89 1.89 0 0 1 2 2v12a1.89 1.89 0 0 1-2 2zM4 14v12h12V14z" fill="currentColor"/><path d="M22 19h-2v-9H10V8h10a1.89 1.89 0 0 1 2 2z" fill="currentColor"/><path d="M26 14h-2V6h-8V4h8a1.89 1.89 0 0 1 2 2z" fill="currentColor"/><path d="M24 17v2h2.8L22 24.4V22h-2v6h6v-2h-2.8l4.8-5.4V23h2v-6h-6z" fill="currentColor"/>`),
		g.Group(children),
	)
}