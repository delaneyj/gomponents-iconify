package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M2.5 0v2.5l-.75-.75L1 2.5L2.5 4L1 5.5l.75.75l.75-.75V8H3l3.5-2.5l-2.25-1.53L6.5 2.5L3 0h-.5zm1 1.5l1.5 1l-1.5 1v-2zm0 3l1.5 1l-1.5 1v-2z"/>`),
		g.Group(children),
	)
}