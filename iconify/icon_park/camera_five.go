package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><circle cx="24" cy="21" r="16" fill="#2F88FF" stroke="#000"/><circle cx="24" cy="21" r="7" fill="#43CCF8" stroke="#fff"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M16 43L32 43"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M24 37V43"/></g>`),
		g.Group(children),
	)
}