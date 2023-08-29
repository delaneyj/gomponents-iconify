package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23.555 6.88L17.328.651L3.979 14.002l-.565 2.826l3.965 3.966l2.827-.566L23.555 6.88Zm-2.829 0L9.22 18.385l-1.184.236l-2.451-2.451l.236-1.184L17.328 3.481l3.398 3.398ZM1.386 19.612l3.208 3.208l1.414-1.414L2.8 18.2l-1.414 1.414Z"/>`),
		g.Group(children),
	)
}