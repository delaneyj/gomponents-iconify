package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraDomeFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 7.25A3.25 3.25 0 0 1 7.25 4h33.5A3.25 3.25 0 0 1 44 7.25v2.5a3.25 3.25 0 0 1-2 3V26c0 9.941-8.059 18-18 18S6 35.941 6 26V12.75a3.25 3.25 0 0 1-2-3v-2.5ZM8.5 13v13c0 8.56 6.94 15.5 15.5 15.5c8.56 0 15.5-6.94 15.5-15.5V13h-31Zm33-3.25v-2.5a.75.75 0 0 0-.75-.75H7.25a.75.75 0 0 0-.75.75v2.5c0 .414.336.75.75.75h33.5a.75.75 0 0 0 .75-.75ZM15.5 27a8.5 8.5 0 1 1 17 0a8.5 8.5 0 0 1-17 0ZM24 16c-6.075 0-11 4.925-11 11s4.925 11 11 11s11-4.925 11-11s-4.925-11-11-11Zm-6 11a6 6 0 1 1 12 0a6 6 0 0 1-12 0Z"/>`),
		g.Group(children),
	)
}