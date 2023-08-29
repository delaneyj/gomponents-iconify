package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quantcast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.146 27.297A13.366 13.366 0 0 0 26.318 16c0-7.365-5.896-13.333-13.161-13.333S.001 8.636.001 16c0 7.365 5.891 13.333 13.156 13.333h18.844v-2.036z"/>`),
		g.Group(children),
	)
}