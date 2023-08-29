package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanPersonTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6a3 3 0 0 1 3-3h1.5a.5.5 0 0 1 0 1H6a2 2 0 0 0-2 2v1.5a.5.5 0 0 1-1 0V6Zm9-2.5a.5.5 0 0 1 .5-.5H14a3 3 0 0 1 3 3v1.5a.5.5 0 0 1-1 0V6a2 2 0 0 0-2-2h-1.5a.5.5 0 0 1-.5-.5ZM3.5 12a.5.5 0 0 1 .5.5V14a2 2 0 0 0 1.098 1.785c.27.138.577.215.902.215h1.5a.5.5 0 0 1 0 1H6a3 3 0 0 1-3-3v-1.5a.5.5 0 0 1 .5-.5Zm13 0a.5.5 0 0 1 .5.5V14a3 3 0 0 1-3 3h-1.5a.5.5 0 0 1 0-1H14a2 2 0 0 0 2-2v-1.5a.5.5 0 0 1 .5-.5Zm-4 3a1.5 1.5 0 0 0-1.415 1h-2.17A1.5 1.5 0 0 0 7.5 15H6a1 1 0 0 1-.975-.776A1.5 1.5 0 0 1 6.5 13h7a1.5 1.5 0 0 1 1.475 1.225A1 1 0 0 1 14 15h-1.5Zm.25-5.75a2.75 2.75 0 1 0-5.5 0a2.75 2.75 0 0 0 5.5 0Z"/>`),
		g.Group(children),
	)
}