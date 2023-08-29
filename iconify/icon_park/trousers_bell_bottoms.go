package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrousersBellBottoms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M16 12L18 4H30L32 12V24L36 41L24 44L12 41L16 24V12Z"/><path stroke="#fff" d="M24 44V16"/><path stroke="#000" d="M12 41L24 44L36 41"/></g>`),
		g.Group(children),
	)
}