package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubscriptionManagement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19h5v4l-2.5-1.5L6 23v-4zM20 1h-8v22l2.5-1.5L17 23v-2h3a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2Zm-6 4h2v2h-2Zm0 5h2v2h-2Zm0 5h2v2h-2ZM4 1a2.006 2.006 0 0 0-2 2v16a2.006 2.006 0 0 0 2 2V3h7V1Z"/><path fill="currentColor" d="M6 5h2v2H6zm0 5h2v2H6zm0 5h2v2H6z"/>`),
		g.Group(children),
	)
}