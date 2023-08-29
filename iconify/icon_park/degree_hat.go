package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DegreeHat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M2 17.4L23.0222 9L44.0444 17.4L23.0222 25.8L2 17.4Z"/><path stroke-linecap="round" d="M44.0444 17.51V26.7332"/><path stroke-linecap="round" d="M11.5557 21.8252V34.2666C11.5557 34.2666 16.3658 38.9999 23.0224 38.9999C29.679 38.9999 34.4891 34.2666 34.4891 34.2666V21.8252"/></g>`),
		g.Group(children),
	)
}