package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 3H3v18h18V11h-2v8H5V5h8V3zm-6 8h2v6H7v-6zm6-4h-2v10h2V7zm2 6h2v4h-2v-4zm6-8h-6v2h6V5z"/>`),
		g.Group(children),
	)
}