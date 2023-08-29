package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirtelXstream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.112 42.5h-5.725L19.87 27.985l-5.71 5.708H8.585l8.563-8.562L5.71 13.695h5.766l8.542 8.54L36.756 5.5h5.534L22.695 25.093Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.017 22.189l2.968 2.968l-3.08 3.08l-3.036-3.036Z"/>`),
		g.Group(children),
	)
}