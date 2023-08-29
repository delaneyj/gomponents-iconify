package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mosque(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.004 2.002l-.002 1.081A6.005 6.005 0 0 1 17.917 8H22v2h-1v12H3V10H2V8h4.083a6.005 6.005 0 0 1 4.919-4.917l.002-1.085l2 .004ZM8.126 8h7.748a4.002 4.002 0 0 0-7.748 0ZM5 10v10h3v-6h8v6h3V10H5Zm9 10v-4h-4v4h4Z"/>`),
		g.Group(children),
	)
}