package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenRotation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 24C4 35.0457 12.9543 44 24 44L19 39"/><path d="M44 24C44 12.9543 35.0457 4 24 4L29 9"/><path fill="#2F88FF" d="M30 41L7 18L18 7L41 30L30 41Z"/></g>`),
		g.Group(children),
	)
}