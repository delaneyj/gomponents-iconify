package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTreemapTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 3a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6Zm1 1v12H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h1Zm1 12v-3h8v1a2 2 0 0 1-2 2H8Zm8-4H8V4h6a2 2 0 0 1 2 2v6Z"/>`),
		g.Group(children),
	)
}