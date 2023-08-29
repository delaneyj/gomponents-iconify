package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.84 22.73L18.11 20H3v-1l2-2v-6c0-1.14.29-2.27.83-3.28L1.11 3l1.28-1.27l19.72 19.73l-1.27 1.27M19 15.8V11c0-3.1-2.03-5.83-5-6.71V4a2 2 0 0 0-2-2a2 2 0 0 0-2 2v.29c-.61.18-1.2.45-1.74.8L19 15.8M12 23a2 2 0 0 0 2-2h-4a2 2 0 0 0 2 2Z"/>`),
		g.Group(children),
	)
}