package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 3.25a1.25 1.25 0 1 0-2.5 0V6.5H3.25a1.25 1.25 0 1 0 0 2.5H6.5v11.5a5 5 0 0 0 5 5H23v3.25a1.25 1.25 0 1 0 2.5 0V25.5h3.25a1.25 1.25 0 1 0 0-2.5H11.5A2.5 2.5 0 0 1 9 20.5V3.25Zm14 8.25v10h2.5v-10a5 5 0 0 0-5-5h-10V9h10a2.5 2.5 0 0 1 2.5 2.5Z"/>`),
		g.Group(children),
	)
}