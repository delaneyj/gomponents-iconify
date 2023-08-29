package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.4 12.5L12 18.38L11 19v-5.62l-8 5L2 19V6l9 5.62V6l10.4 6.5m-1.9 0L12 7.8v9.4l7.5-4.7m-9 0L3 7.8v9.4l7.5-4.7Z"/>`),
		g.Group(children),
	)
}