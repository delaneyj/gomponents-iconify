package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseRemind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 38C42 38 36 33 36 19C36 12.3726 30.6274 7 24 7C21.46 7 19.1042 7.78918 17.1647 9.13571M30 38H6C6 38 11.5692 33.359 11.9765 20.5"/><path stroke="#000" stroke-width="4" d="M18 38L30 38C30 41.3137 27.3137 44 24 44C20.6863 44 18 41.3137 18 38Z"/><path fill="#000" fill-rule="evenodd" d="M24 2C21.7909 2 20 3.79086 20 6H28C28 3.79086 26.2091 2 24 2Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7 6.5L41 44.5"/></g>`),
		g.Group(children),
	)
}