package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noview(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11.478 17.568A4.752 4.752 0 0 1 11.193 16A4.814 4.814 0 0 1 16 11.193c.552 0 1.074.113 1.568.285l2.283-2.283c-1.31-.548-2.623-.91-3.85-.91C8.454 8.286 2.5 16 2.5 16s2.167 2.79 5.53 5.017l3.448-3.45zm12.04-6.383l-3.056 3.056c.217.547.345 1.14.345 1.76A4.815 4.815 0 0 1 16 20.809a4.757 4.757 0 0 1-1.76-.345l-2.47 2.47c1.328.48 2.746.783 4.23.783c5.77 0 13.5-7.715 13.5-7.715s-2.64-2.626-5.982-4.815zm2.024-6.268L4.855 25.604L6.27 27.02L26.956 6.332l-1.414-1.415z"/>`),
		g.Group(children),
	)
}