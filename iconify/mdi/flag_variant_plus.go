package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagVariantPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 3c.6 0 1 .4 1 1v.9c1.1-.5 2.5-.9 4-.9c3 0 3 2 5 2c3 0 4-2 4-2v8s-1 2-4 2s-3-2-5-2c-3 0-4 2-4 2v7H5V4c0-.6.4-1 1-1m12 12v3h-3v2h3v3h2v-3h3v-2h-3v-3h-2Z"/>`),
		g.Group(children),
	)
}