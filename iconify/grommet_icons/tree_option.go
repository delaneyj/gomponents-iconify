package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeOption(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m7 13l4.375-7H9l3-4l3 4h-2.375L17 13h-2l4 6.667H5L9 13H7Zm5 11v-4"/>`),
		g.Group(children),
	)
}