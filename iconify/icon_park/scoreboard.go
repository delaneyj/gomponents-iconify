package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scoreboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="40" height="28" x="4" y="12" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#000" stroke-linecap="round" d="M14 6V12"/><path stroke="#000" stroke-linecap="round" d="M34 6V12"/><path stroke="#fff" stroke-linecap="round" d="M10.2266 24L15.0006 19.0166V33"/><path stroke="#fff" stroke-linecap="round" d="M24 12V40"/><ellipse cx="34" cy="26" stroke="#fff" rx="3" ry="7"/><path stroke="#000" stroke-linecap="round" d="M21 12H27"/><path stroke="#000" stroke-linecap="round" d="M21 40H27"/></g>`),
		g.Group(children),
	)
}