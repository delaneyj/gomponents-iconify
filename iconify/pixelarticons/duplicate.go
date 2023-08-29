package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Duplicate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 3h12v4h4v14H7v-4H3V3h2zm10 4V5H5v10h2V7h8zM9 17v2h10V9H9v8z"/>`),
		g.Group(children),
	)
}