package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4h11a3 3 0 0 1 3 3v11a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3m0 1a2 2 0 0 0-2 2v8h5v.5a2.5 2.5 0 0 0 2.5 2.5a2.5 2.5 0 0 0 2.5-2.5V15h5V7a2 2 0 0 0-2-2H6M4 18a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2v-2h-4.04c-.24 1.7-1.7 3-3.46 3s-3.22-1.3-3.46-3H4v2Z"/>`),
		g.Group(children),
	)
}