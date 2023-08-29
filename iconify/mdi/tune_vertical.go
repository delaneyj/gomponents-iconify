package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuneVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 3H5v6h2V3m12 0h-2v10h2V3M3 13h2v8h2v-8h2v-2H3v2m12-6h-2V3h-2v4H9v2h6V7m-4 14h2V11h-2v10m4-6v2h2v4h2v-4h2v-2h-6Z"/>`),
		g.Group(children),
	)
}