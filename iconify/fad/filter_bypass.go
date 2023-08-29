package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterBypass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M25 67.995A4.002 4.002 0 0 1 29.004 64H227a4 4 0 0 1 4 3.995v10.01A3.99 3.99 0 0 1 227 82H29a4 4 0 0 1-4-3.995v-10.01z"/>`),
		g.Group(children),
	)
}