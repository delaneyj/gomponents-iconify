package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 5.586L9.914 12L3.5 18.414L2.086 17l5-5l-5-5L3.5 5.586ZM12 18h10v2H12v-2Z"/>`),
		g.Group(children),
	)
}