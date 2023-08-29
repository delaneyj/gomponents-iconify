package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alien(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 17a2.5 2.5 0 0 0 2 0"/><path d="M12 3C7.336 3 4.604 5.331 4.138 8.595a11.816 11.816 0 0 0 2 8.592a10.777 10.777 0 0 0 3.199 3.064c1.666 1 3.664 1 5.33 0a10.777 10.777 0 0 0 3.199-3.064a11.89 11.89 0 0 0 2-8.592C19.4 5.33 16.668 3 12.004 3zm-4 8l2 2m6-2l-2 2"/></g>`),
		g.Group(children),
	)
}