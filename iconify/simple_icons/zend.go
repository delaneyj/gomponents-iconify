package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 .467L9.01 20.217h11.53a3.46 3.46 0 0 0 3.46-3.46V.468zM3.459 3.783C1.585 3.783 0 5.153 0 7.1v16.433l15.063-19.75H3.459Z"/>`),
		g.Group(children),
	)
}