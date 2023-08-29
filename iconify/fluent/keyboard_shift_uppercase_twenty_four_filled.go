package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardShiftUppercaseTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.32 2.603a1.75 1.75 0 0 0-2.643 0l-8.245 9.504c-.983 1.133-.179 2.897 1.321 2.897H7v2.247c0 .966.784 1.75 1.75 1.75h6.5A1.75 1.75 0 0 0 17 17.25v-2.247h3.245c1.5 0 2.305-1.764 1.322-2.897l-8.245-9.504ZM7.75 20.5a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5h-8.5Z"/>`),
		g.Group(children),
	)
}