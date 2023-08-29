package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosFaceRecognition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 34v10h10m20 0h10V34M34 4h10v10M14 4H4v10m12 20s3 3 8 3s8-3 8-3m-8-20v9c0 2-2 4-4 4h-1m15-13v2m-20-2v2"/>`),
		g.Group(children),
	)
}