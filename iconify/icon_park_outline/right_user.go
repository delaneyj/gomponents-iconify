package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="12" r="8"/><path d="M42 44c0-9.941-8.059-18-18-18S6 34.059 6 44"/><path d="m30 36l-8 8l-4-4"/></g>`),
		g.Group(children),
	)
}