package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtractSelectionOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M32 16H41C42.1046 16 43 16.8954 43 18V41C43 42.1046 42.1046 43 41 43H18C16.8954 43 16 42.1046 16 41V32"/><rect width="27" height="27" x="5" y="5" rx="2"/><path d="M18 5L5 19"/><path d="M27 5L5 29"/><path d="M32 10L12 32"/><path d="M32 21L22 32"/></g>`),
		g.Group(children),
	)
}