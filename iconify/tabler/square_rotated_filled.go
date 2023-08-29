package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareRotatedFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m9.793 2.893l-6.9 6.9c-1.172 1.171-1.172 3.243 0 4.414l6.9 6.9c1.171 1.172 3.243 1.172 4.414 0l6.9-6.9c1.172-1.171 1.172-3.243 0-4.414l-6.9-6.9c-1.171-1.172-3.243-1.172-4.414 0z"/></g>`),
		g.Group(children),
	)
}