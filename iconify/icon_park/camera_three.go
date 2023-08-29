package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><rect width="36" height="26" x="6" y="14" stroke-linejoin="round" rx="2"/><path fill="#2F88FF" stroke-linejoin="round" d="M10 14L12.1667 8H19.8333L22 14H10Z"/><circle cx="29" cy="27" r="7" fill="#2F88FF" stroke-linejoin="round"/><path d="M36 10V14"/></g>`),
		g.Group(children),
	)
}