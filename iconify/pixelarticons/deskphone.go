package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deskphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18v18H3V3zm2 2v6h8V5H5zm10 0v14h4V5h-4zm-2 14v-2h-3v2h3zm-5 0v-2H5v2h3zm-3-4h3v-2H5v2zm5-2v2h3v-2h-3z"/>`),
		g.Group(children),
	)
}