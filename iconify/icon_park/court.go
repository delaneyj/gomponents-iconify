package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Court(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 36H44V12H4V36H24ZM24 36V28M24 12V20"/><circle cx="24" cy="24" r="4"/><path d="M11 24C11 26.2091 9.20914 28 7 28H4V20H7C9.20914 20 11 21.7909 11 24Z"/><path d="M37 24C37 26.2091 38.7909 28 41 28H44V20H41C38.7909 20 37 21.7909 37 24Z"/></g>`),
		g.Group(children),
	)
}