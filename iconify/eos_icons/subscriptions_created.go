package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubscriptionsCreated(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.01 6.01h16v2h-16zm2-4h12v2h-12zM20 10H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2Zm-9.7 10L7 16.76l1.4-1.4l1.9 1.9l5.3-5.3l1.4 1.4Z"/>`),
		g.Group(children),
	)
}