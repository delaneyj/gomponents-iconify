package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.879 18.123a3 3 0 0 1 0-4.243L13.88 3.879a3 3 0 0 1 4.243 0L28.124 13.88a3 3 0 0 1 0 4.243L18.123 28.124a3 3 0 0 1-4.243 0L3.879 18.123Zm1.414-2.829a1 1 0 0 0 0 1.415l10.001 10a1 1 0 0 0 1.414 0l10.002-10a1 1 0 0 0 0-1.415L16.708 5.293a1 1 0 0 0-1.414 0L5.293 15.294Z"/>`),
		g.Group(children),
	)
}