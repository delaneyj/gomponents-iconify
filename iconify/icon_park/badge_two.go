package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BadgeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M16 43.9999L24 39.9999L32 43.9999V24.9443C29.877 26.8445 27.0734 27.9999 24 27.9999C20.9266 27.9999 18.123 26.8445 16 24.9443V43.9999Z"/><path fill="#2F88FF" stroke="#000" d="M36 16C36 19.554 34.455 22.7471 32 24.9444C29.877 26.8446 27.0734 28 24 28C20.9266 28 18.123 26.8446 16 24.9444C13.545 22.7471 12 19.554 12 16C12 9.37258 17.3726 4 24 4C30.6274 4 36 9.37258 36 16Z"/><path stroke="#fff" d="M24 21V11L22 12M24 21H26M24 21H22"/></g>`),
		g.Group(children),
	)
}