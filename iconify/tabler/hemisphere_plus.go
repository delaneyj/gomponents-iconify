package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HemispherePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 9a9 3 0 1 0 18 0A9 3 0 1 0 3 9"/><path d="M3 9a9 9 0 0 0 9 9m8.396-5.752A8.978 8.978 0 0 0 21 9m-5 10h6m-3-3v6"/></g>`),
		g.Group(children),
	)
}