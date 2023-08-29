package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cicsplex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 18.6V17H6v1.6a7.383 7.383 0 0 0 7.4 7.4H15v-2h-1.6A5.378 5.378 0 0 1 8 18.6zm20-.6h-8a2.006 2.006 0 0 0-2 2v8a2.006 2.006 0 0 0 2 2h8a2.006 2.006 0 0 0 2-2v-8a2.006 2.006 0 0 0-2-2zm-8 10v-8h8v8zm4-14.6V15h2v-1.6A7.383 7.383 0 0 0 18.6 6H17v2h1.6a5.378 5.378 0 0 1 5.4 5.4zM12 2H4a2.006 2.006 0 0 0-2 2v8a2.006 2.006 0 0 0 2 2h8a2.006 2.006 0 0 0 2-2V4a2.006 2.006 0 0 0-2-2zM4 12V4h8v8z"/>`),
		g.Group(children),
	)
}