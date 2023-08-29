package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedalOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M6 6L16 18"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M42 6L32 18"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M31 4L26 16"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M17 4L22 16"/><circle cx="24" cy="30" r="14" fill="#2F88FF" stroke="#000"/><circle cx="24" cy="30" r="7" fill="#43CCF8" stroke="#fff"/></g>`),
		g.Group(children),
	)
}