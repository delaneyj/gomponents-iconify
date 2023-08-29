package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditFlipVerticalTwoArrowDesignFlipReflectUpDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11 13.5l-4-4l-4 4h8zm0-13l-4 4l-4-4h8zM13.5 7h-1m-2 0h-1m-2 0h-1m-2 0h-1m-2 0h-1"/>`),
		g.Group(children),
	)
}