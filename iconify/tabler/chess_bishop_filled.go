package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessBishopFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 2a2 2 0 0 1 1.386 3.442c.646.28 1.226.62 1.74 1.017l-3.833 3.834l-.083.094a1 1 0 0 0 1.403 1.403l.094-.083l3.814-3.813C17.498 9.244 18 10.964 18 13c0 1.913-1.178 3.722-3.089 3.973l-.2.02L14.5 17h-5C7.374 17 6 15.076 6 13c0-3.68 1.57-6.255 4.613-7.56A2 2 0 0 1 12 2zm0 3v1m6 12H6a1 1 0 0 0-1 1a2 2 0 0 0 2 2h10a2 2 0 0 0 1.987-1.768l.011-.174A1 1 0 0 0 18 18z"/></g>`),
		g.Group(children),
	)
}