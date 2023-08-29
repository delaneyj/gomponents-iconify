package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficEvents(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1.914l3.643 8.687l.406.899l-.025.01L19.165 19H21v2H3v-2h1.835l3.14-7.49l-.024-.01l.406-.9L12 1.915ZM10.149 11.5l-1.145 2.73h5.992l-1.145-2.73H10.15Zm2.863-2L12 7.086L10.988 9.5h2.024Zm2.823 6.73h-7.67L7.004 19h9.992l-1.161-2.77Z"/>`),
		g.Group(children),
	)
}