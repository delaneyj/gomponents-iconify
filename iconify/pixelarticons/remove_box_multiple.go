package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveBoxMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 3H3v14h14V3H5zm10 2v10H5V5h10zm4 2v12H7v2h14V7h-2zm-6 2H7v2h6V9z"/>`),
		g.Group(children),
	)
}