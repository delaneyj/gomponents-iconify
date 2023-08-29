package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M41 5.5H7V33.5H41V5.5Z"/><path stroke="#000" stroke-linecap="round" d="M16 41.5L24 33.5L32 41.5"/><path stroke="#fff" stroke-linecap="round" d="M13.9239 24.6628L19.5664 19.155L24.0083 23.4999L33.9669 13.5208"/><path stroke="#000" stroke-linecap="round" d="M4 5.5H44"/></g>`),
		g.Group(children),
	)
}