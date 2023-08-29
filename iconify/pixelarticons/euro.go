package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Euro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 4h10v2H9v3h7v2H9v2h7v2H9v3h10v2H7v-5H5v-2h2v-2H5V9h2V4h2z"/>`),
		g.Group(children),
	)
}