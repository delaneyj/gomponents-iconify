package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraSwitchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 4h-3.2L15 2H9L7.2 4H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2M9.9 4h4.2l1.8 2H20v12H4V6h4.1m6.9 5H9V8.5L5.5 12L9 15.5V13h6v2.5l3.5-3.5L15 8.5V11Z"/>`),
		g.Group(children),
	)
}