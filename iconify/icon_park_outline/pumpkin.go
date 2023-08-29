package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pumpkin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="40" height="26" x="4" y="14" rx="13"/><ellipse cx="24" cy="27" rx="8" ry="13"/><path stroke-linecap="round" stroke-linejoin="round" d="M30 6h-3a3 3 0 0 0-3 3v5"/></g>`),
		g.Group(children),
	)
}