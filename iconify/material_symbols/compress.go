package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 14v-2h16v2H4Zm0-3V9h16v2H4Zm7 11v-3.2l-1.6 1.6L8 19l4-4l4 4l-1.4 1.4l-1.6-1.55V22h-2Zm1-14L8 4l1.4-1.4L11 4.2V1h2v3.2l1.6-1.6L16 4l-4 4Z"/>`),
		g.Group(children),
	)
}