package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChildrenPyramid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="18" height="8" x="15" y="17" fill="#2F88FF"/><path fill="#2F88FF" d="M10 25H38V33H10V25Z"/><path fill="#2F88FF" d="M5 33H43V41H5V33Z"/><path d="M24 17V7"/><path d="M29 7L19 7"/></g>`),
		g.Group(children),
	)
}