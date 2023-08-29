package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCardOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.586 1H21v22H3V5.586L7.586 1Zm.828 2L5 6.414V21h14V3H8.414ZM10.5 9.006h3v7.982h-2v-5.982h-1v-2Z"/>`),
		g.Group(children),
	)
}