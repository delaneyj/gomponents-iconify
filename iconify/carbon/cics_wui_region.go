package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CicsWuiRegion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 22h2v6h-2v2h6v-2h-2v-6h2v-2h-6v2zm-4 6h-2v-8h-2v8.6a1.453 1.453 0 0 0 1.5 1.4h3a1.453 1.453 0 0 0 1.5-1.4V20h-2zm-7.8-8l-.2 8.5L9 22H7l-1 6.5l-.2-8.5H4l.72 10H7l1-6.5L9 30h2.28L12 20h-1.8zm5.8-9h-3V8h-2v3H8v2h3v3h2v-3h3v-2z"/><path fill="currentColor" d="M26 4H6a2.006 2.006 0 0 0-2 2v12h2V6h20v12h2V6a2.006 2.006 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}