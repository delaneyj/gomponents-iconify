package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sketch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path fill="#43CCF8" stroke="#fff" stroke-linecap="round" d="M18.6 16H29.4L33 20.7059L24 32L15 20.7059L18.6 16Z"/></g>`),
		g.Group(children),
	)
}