package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistancingCorrectFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M5.089 19.288a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25Zm4.224 3.75a4.474 4.474 0 0 0-8.449 0m18.026-17a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25Zm4.224 3.75a4.474 4.474 0 0 0-8.448 0"/><path d="M11.614 19.913a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m3-2.25a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75m3-2.25a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/><path stroke-linecap="round" stroke-linejoin="round" d="M5.386 9.962a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m3.586 5.912l.974.974a.449.449 0 0 0 .684-.056l1.942-2.718"/></g>`),
		g.Group(children),
	)
}