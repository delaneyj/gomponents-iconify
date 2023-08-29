package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><rect width="36" height="26" x="6" y="14" stroke-linejoin="round" rx="2"/><path stroke-linejoin="round" d="m10 14l2.167-6h7.666L22 14H10Z"/><circle cx="29" cy="27" r="7" stroke-linejoin="round"/><path d="M36 10v4"/></g>`),
		g.Group(children),
	)
}