package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookBookmarkLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 26H72a30 30 0 0 0-30 30v168a6 6 0 0 0 6 6h144a6 6 0 0 0 0-12H54v-2a18 18 0 0 1 18-18h136a6 6 0 0 0 6-6V32a6 6 0 0 0-6-6Zm-90 12h52v78l-22.41-16.8a6 6 0 0 0-7.2 0L118 116Zm84 148H72a29.87 29.87 0 0 0-18 6V56a18 18 0 0 1 18-18h34v90a6 6 0 0 0 9.6 4.8l28.4-21.3l28.41 21.3A6 6 0 0 0 182 128V38h20Z"/>`),
		g.Group(children),
	)
}