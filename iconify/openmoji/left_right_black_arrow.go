package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftRightBlackArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3f3f3f" d="m48.341 19.828l-3.7 4.1l9.2 8.6H17.54l9.2-8.6l-3.7-4.1l-16.7 15.8l16.7 15.8l3.7-4.1l-9.2-8.6h36.301l-9.2 8.6l3.7 4.1l16.7-15.8l-16.7-15.8z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m48.341 19.828l-3.7 4.1l9.2 8.6H17.54l9.2-8.6l-3.7-4.1l-16.7 15.8l16.7 15.8l3.7-4.1l-9.2-8.6h36.301l-9.2 8.6l3.7 4.1l16.7-15.8l-16.7-15.8z"/>`),
		g.Group(children),
	)
}