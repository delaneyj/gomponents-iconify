package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Douban(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M20 6H4V4h16v2m0 12v2H4v-2h3.33l-1.07-4H5V8h14v6h-1.26l-1.07 4H20M7 12h10v-2H7v2m2.4 6h5.2l1.07-4H8.33l1.07 4z" fill="currentColor"/>`),
		g.Group(children),
	)
}