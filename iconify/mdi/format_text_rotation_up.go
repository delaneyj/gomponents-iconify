package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatTextRotationUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 12v1.5l11 4.75v-2.1l-2.2-.9v-5l2.2-.9v-2.1L3 12m7 2.62l-5-1.87l5-1.87v3.74m8-10.37l-3 3h2v12.5h2V7.25h2l-3-3Z"/>`),
		g.Group(children),
	)
}