package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Incoming(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" stroke-linejoin="round" d="M18 13C18 10.7909 16.2091 9 14 9C11.7909 9 10 10.7909 10 13C10 15.2091 11.7909 17 14 17C16.2091 17 18 15.2091 18 13Z"/><path fill="#2F88FF" stroke-linejoin="round" d="M37 13C37 10.7909 35.2091 9 33 9C30.7909 9 29 10.7909 29 13C29 15.2091 30.7909 17 33 17C35.2091 17 37 15.2091 37 13Z"/><path stroke-linecap="round" d="M4 13H10"/><path stroke-linecap="round" d="M18 13H29"/><path stroke-linecap="round" d="M37 13H44"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 29L4 35L10 41"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 35H44"/></g>`),
		g.Group(children),
	)
}