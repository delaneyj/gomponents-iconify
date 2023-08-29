package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="14" height="27" x="17" y="4" fill="#2F88FF" rx="7"/><path stroke-linecap="round" d="M9 23C9 31.2843 15.7157 38 24 38C32.2843 38 39 31.2843 39 23"/><path stroke-linecap="round" d="M24 38V44"/></g>`),
		g.Group(children),
	)
}