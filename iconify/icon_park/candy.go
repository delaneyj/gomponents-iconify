package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Candy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><circle cx="24" cy="24" r="10" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round"/><path stroke="#fff" stroke-linecap="round" d="M24 28C21.7909 28 20 26.2091 20 24"/><path stroke="#000" stroke-linejoin="round" d="M16.6875 16.8125L3.90824 14.9668L14.8418 4.03324L16.6875 16.8125Z"/><path stroke="#000" stroke-linejoin="round" d="M31.3125 31.3125L44.0918 33.1582L33.1582 44.0918L31.3125 31.3125Z"/></g>`),
		g.Group(children),
	)
}