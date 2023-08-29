package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="40" height="28" x="4" y="13" stroke="currentColor" stroke-linejoin="round" stroke-width="4" rx="14"/><circle cx="31" cy="22" r="2" fill="currentColor"/><circle cx="35" cy="27" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 27h10m-10 0h10m2-20v6m0-6v6m-7 9v10"/></g>`),
		g.Group(children),
	)
}