package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 7V1H19.5L23 4.5V23h-3M18 1v5h5M3 11h13v12H3V11Zm4 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-2 7l2-3l2 2l4-6l3 4"/>`),
		g.Group(children),
	)
}