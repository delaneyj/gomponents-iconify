package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 6H4v12.01h16V11h-5z" opacity=".3"/><path fill="currentColor" d="M4 4c-1.1 0-2 .9-2 2v12.01c0 1.1.9 1.99 2 1.99h16c1.1 0 2-.9 2-2v-8l-6-6H4zm16 14.01H4V6h11v5h5v7.01z"/>`),
		g.Group(children),
	)
}