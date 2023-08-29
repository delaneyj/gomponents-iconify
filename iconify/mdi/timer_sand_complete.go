package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerSandComplete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 22H6v-6l4-4l-4-4V2h12v6l-4 4l4 4M8 7.5l4 4l4-4V4H8m4 8.5l-4 4V20h8v-3.5M14 18h-4v-.8l2-2l2 2Z"/>`),
		g.Group(children),
	)
}