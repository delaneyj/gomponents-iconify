package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Consignment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M8 14C8 12.8954 8.89543 12 10 12H38C39.1046 12 40 12.8954 40 14V30C40 31.1046 39.1046 32 38 32H10C8.89543 32 8 31.1046 8 30V14Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M16 12V32"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M32 12V32"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M28 12L36 12"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M12 12L20 12"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M12 32L20 32"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M28 32L36 32"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M4 38L44 38"/><path stroke="#000" stroke-linecap="round" d="M18 38V40"/><path stroke="#000" stroke-linecap="round" d="M12 38V40"/><path stroke="#000" stroke-linecap="round" d="M6 38V40"/><path stroke="#000" stroke-linecap="round" d="M24 38V40"/><path stroke="#000" stroke-linecap="round" d="M30 38V40"/><path stroke="#000" stroke-linecap="round" d="M36 38V40"/><path stroke="#000" stroke-linecap="round" d="M42 38V40"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M30 12V6H18V12"/></g>`),
		g.Group(children),
	)
}