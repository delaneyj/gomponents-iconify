package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuneVerticalVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 12.14V2H6v10.14c-1.72.45-3 2-3 3.86s1.28 3.41 3 3.86V22h2v-2.14c1.72-.45 3-2 3-3.86s-1.28-3.41-3-3.86M7 14c1.1 0 2 .9 2 2s-.9 2-2 2s-2-.9-2-2s.9-2 2-2M18 2h-2v2.14c-1.72.45-3 2-3 3.86s1.28 3.41 3 3.86V22h2V11.86c1.72-.45 3-2 3-3.86s-1.28-3.41-3-3.86V2m-1 4c1.1 0 2 .9 2 2s-.9 2-2 2s-2-.9-2-2s.9-2 2-2Z"/>`),
		g.Group(children),
	)
}