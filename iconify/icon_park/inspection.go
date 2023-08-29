package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inspection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M43 33V19H5V41C5 42.1046 5.89543 43 7 43H24"/><path stroke-linejoin="round" d="M5 10C5 8.89543 5.89543 8 7 8H41C42.1046 8 43 8.89543 43 10V19H5V10Z"/><path stroke-linecap="round" d="M16 5V13"/><path stroke-linecap="round" d="M32 5V13"/><circle cx="30" cy="32" r="7" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M36 37L42 42"/></g>`),
		g.Group(children),
	)
}