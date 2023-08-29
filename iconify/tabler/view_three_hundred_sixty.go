package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewThreeHundredSixty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/><path d="M8 12a4 9 0 1 0 8 0a4 9 0 1 0-8 0"/><path d="M3 12c0 2.21 4.03 4 9 4s9-1.79 9-4s-4.03-4-9-4s-9 1.79-9 4z"/></g>`),
		g.Group(children),
	)
}