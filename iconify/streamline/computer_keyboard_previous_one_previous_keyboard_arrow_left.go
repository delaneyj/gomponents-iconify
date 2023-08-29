package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerKeyboardPreviousOnePreviousKeyboardArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 7h-10M7 3.5L3.5 7L7 10.5m-6.5-7v7"/>`),
		g.Group(children),
	)
}