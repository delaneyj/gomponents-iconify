package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 9v0a4 4 0 0 0-4 4v6a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-6a4 4 0 0 0-4-4v0M7 9h10M7 9V6m10 3V6m-5 0v3m0-6h.01M7 3h.01M17 3h.01"/><path d="M3 13c0 1 .6 3 3 3s3-2 3-3c0 1 .6 3 3 3s3-2 3-3c0 1 .6 3 3 3s3-2 3-3"/></g>`),
		g.Group(children),
	)
}