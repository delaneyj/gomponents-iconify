package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileBlocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11h-2V9h-6V3H5v18h6v2H3V1Zm12 2.414V7h3.586L15 3.414ZM18 14.5a3.5 3.5 0 0 0-3.08 5.165l4.745-4.744A3.483 3.483 0 0 0 18 14.5Zm3.08 1.835l-4.745 4.744a3.5 3.5 0 0 0 4.745-4.745ZM12.5 18a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Z"/>`),
		g.Group(children),
	)
}