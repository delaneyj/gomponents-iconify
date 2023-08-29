package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemDevice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.533 1H21v21H2V6.588h4.533V1Zm2 5.588h4.534V20H19V3H8.533v3.588ZM11.067 20V8.588H4V20h7.067ZM6 16h3v2H6v-2Z"/>`),
		g.Group(children),
	)
}