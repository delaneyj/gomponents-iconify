package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SleepTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4 12v24m40-7v7m0-7H4m18-13v13h22V19a3 3 0 0 0-3-3H22Z"/><circle cx="13" cy="20" r="3"/></g>`),
		g.Group(children),
	)
}