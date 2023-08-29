package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stadium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 5L3 7V3l4 2m11-2v4l4-2l-4-2m-7-1v4l4-2l-4-2m-6 8c1.4.5 3.8 1 7 1s5.6-.5 7-1c0-.2-2.8-1-7-1s-7 .9-7 1m10 7H9v4.9c-4.1-.4-7-1.5-7-2.9v-9c0-1.7 4.5-3 10-3s10 1.3 10 3v9c0 1.3-2.9 2.5-7 2.9V17Z"/>`),
		g.Group(children),
	)
}