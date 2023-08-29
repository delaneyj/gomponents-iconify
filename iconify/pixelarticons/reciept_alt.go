package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecieptAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 2H3v20h2v-2h2v2h2v-2h2v2h2v-2h2v2h2v-2h2v2h2V2h-2v2h-2V2h-2v2h-2V2h-2v2H9V2H7v2H5V2zm2 2h2v2h2V4h2v2h2V4h2v2h2v12h-2v2h-2v-2h-2v2h-2v-2H9v2H7v-2H5V6h2V4zm0 4h10v2H7V8zm10 4H7v2h10v-2z"/>`),
		g.Group(children),
	)
}