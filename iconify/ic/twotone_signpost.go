package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneSignpost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 6h11.17l1 1l-1 1H6V6zm12 10H6.83l-1-1l1-1H18v2z" opacity=".3"/><path fill="currentColor" d="M13 10h5l3-3l-3-3h-5V2h-2v2H4v6h7v2H6l-3 3l3 3h5v4h2v-4h7v-6h-7v-2zM6 6h11.17l1 1l-1 1H6V6zm12 10H6.83l-1-1l1-1H18v2z"/>`),
		g.Group(children),
	)
}