package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 3a1 1 0 0 0-2 0v4H3a1 1 0 0 0 0 2h4v11.5a4.5 4.5 0 0 0 4.5 4.5H23v4a1 1 0 1 0 2 0v-4h4a1 1 0 1 0 0-2H11.5A2.5 2.5 0 0 1 9 20.5V3Zm14 8.5v10h2v-10A4.5 4.5 0 0 0 20.5 7h-10v2h10a2.5 2.5 0 0 1 2.5 2.5Z"/>`),
		g.Group(children),
	)
}