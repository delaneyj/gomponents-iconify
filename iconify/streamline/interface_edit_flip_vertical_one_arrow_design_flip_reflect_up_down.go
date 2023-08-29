package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditFlipVerticalOneArrowDesignFlipReflectUpDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.25 1.5L7.18 4.43a.27.27 0 0 1-.36 0L3.75 1.5m6.5 11L7.18 9.57a.27.27 0 0 0-.36 0L3.75 12.5M13.5 7h-1m-2 0h-1m-2 0h-1m-2 0h-1m-2 0h-1"/>`),
		g.Group(children),
	)
}