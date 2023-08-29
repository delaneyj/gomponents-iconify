package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="32" height="26" x="8" y="6" stroke="currentColor" stroke-width="4" rx="2"/><circle cx="14" cy="27" r="2" fill="currentColor"/><circle cx="34" cy="27" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 12h20v10H14zm18 20l8 9m-23-9l-9 9"/></g>`),
		g.Group(children),
	)
}