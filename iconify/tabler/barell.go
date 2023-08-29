package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6.222 4h11.556C19.259 6.37 20 9.037 20 12s-.74 5.63-2.222 8H6.222C4.741 17.63 4 14.963 4 12s.74-5.63 2.222-8z"/><path d="M15 4c.667 2.667 1 5.333 1 8s-.333 5.333-1 8M9 4c-.667 2.667-1 5.333-1 8s.333 5.333 1 8m-4.5-4h15m0-8h-15"/></g>`),
		g.Group(children),
	)
}