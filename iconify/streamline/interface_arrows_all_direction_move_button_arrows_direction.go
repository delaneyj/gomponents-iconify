package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsAllDirectionMoveButtonArrowsDirection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5 2.5l2-2l2 2m-4 9l2 2l2-2M7 .5v13M11.5 5l2 2l-2 2m-9-4l-2 2l2 2m11-2H.5"/>`),
		g.Group(children),
	)
}