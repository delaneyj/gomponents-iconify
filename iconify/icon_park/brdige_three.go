package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrdigeThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M8 12V38"/><path stroke-linecap="round" stroke-linejoin="round" d="M40 12V38"/><path d="M8 38C8 38 16 27 24 27C32 27 40 38 40 38"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 27H44"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 19H44"/><path stroke-linecap="round" d="M24 15L24 27"/><path stroke-linecap="round" d="M16 19L16 27"/><path stroke-linecap="round" d="M32 19L32 27"/></g>`),
		g.Group(children),
	)
}