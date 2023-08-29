package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagChevronBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m240.63 116.91l-42.63-64A19.93 19.93 0 0 0 181.32 44H24a12 12 0 0 0-9.88 18.82l45 65.18l-45 65.18A12 12 0 0 0 24 212h157.32a19.93 19.93 0 0 0 16.68-8.91l42.67-64a19.94 19.94 0 0 0-.04-22.18ZM179.18 188H46.87l33.65-48.74a1.63 1.63 0 0 0 .11-.17a19.91 19.91 0 0 0 0-22.18a1.63 1.63 0 0 0-.11-.17L46.87 68h132.31l40 60Z"/>`),
		g.Group(children),
	)
}