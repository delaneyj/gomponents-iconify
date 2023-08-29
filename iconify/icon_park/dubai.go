package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dubai(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" d="M14 4L14 44"/><path d="M14.5 6C14.5 6 28 13 32 22C36 31 33 44 33 44"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 44H44"/><path stroke-linecap="round" d="M10 15H32"/><path stroke-linecap="round" d="M14 22H22"/><path stroke-linecap="round" d="M14 29H26"/><path stroke-linecap="round" d="M14 36H27"/></g>`),
		g.Group(children),
	)
}