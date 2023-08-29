package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TerminalRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 3v18H1V3h22Zm-2 2H3v14h18V5Zm-2 12h-7v-2h7v2Zm-8.93-5l-3.739 3.744l-1.415-1.413L7.244 12L4.916 9.67L6.33 8.255L10.071 12Z"/>`),
		g.Group(children),
	)
}