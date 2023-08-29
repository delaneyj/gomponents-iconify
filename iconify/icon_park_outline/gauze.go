package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gauze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="26" cy="24" r="17"/><circle cx="26" cy="24" r="7"/><path stroke-linecap="round" stroke-linejoin="round" d="M5 41h21"/></g>`),
		g.Group(children),
	)
}