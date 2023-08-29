package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiveG(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 7H3v6h5v2H3v2h7v-6H5V9h5V7zm11 0h-9v10h9v-6h-4v2h2v2h-5V9h7V7z"/>`),
		g.Group(children),
	)
}