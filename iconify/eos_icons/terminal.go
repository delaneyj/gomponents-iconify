package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 5H4a2.006 2.006 0 0 0-2 2v10a2.006 2.006 0 0 0 2 2h16a2.006 2.006 0 0 0 2-2V7a2.006 2.006 0 0 0-2-2ZM6 17l-1.408-1.412L8.17 12L4.594 8.414L6 7l5 5Zm13 0h-7v-2h7Z"/>`),
		g.Group(children),
	)
}