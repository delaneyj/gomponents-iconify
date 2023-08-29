package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WheelchairActive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 6l3-3l7 3v2l-3 3M9 22a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm0-5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm5-5l5 3l-2 6m2-17a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-6 6l4-4m-6 4l4-4"/>`),
		g.Group(children),
	)
}