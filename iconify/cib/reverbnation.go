package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reverbnation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m32 12.432l-12.193-.041L15.964.76l-3.771 11.631l-12.152.041l9.823 7.141L6.057 31.24h.036l9.948-7.183l9.86 7.183l-3.808-11.641L32 12.422z"/>`),
		g.Group(children),
	)
}