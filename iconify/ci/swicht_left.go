package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwichtLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 18h8a6 6 0 0 0 0-12H8a6 6 0 1 0 0 12Z"/><path d="M8 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/></g>`),
		g.Group(children),
	)
}