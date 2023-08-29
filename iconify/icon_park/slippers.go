package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slippers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 29L44 29V35H4L4 29Z"/><path fill="#2F88FF" d="M23.5293 13C20.0293 17 20.0292 25 20.0293 29H40.0293V20.5C40.0293 20.5 28.0001 15 23.5293 13Z"/></g>`),
		g.Group(children),
	)
}