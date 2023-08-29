package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagVariantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 3a1 1 0 0 1 1 1v.88C8.06 4.44 9.5 4 11 4c3 0 3 2 5 2c3 0 4-2 4-2v8s-1 2-4 2s-3-2-5-2c-3 0-4 2-4 2v7H5V4a1 1 0 0 1 1-1m1 4.25v4.25S9 10 11 10s3 2 5 2s2-1 2-1V7.5S17 8 16 8c-2 0-3-2-5-2S7 7.25 7 7.25Z"/>`),
		g.Group(children),
	)
}