package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoctorTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 3.25A1.75 1.75 0 0 0 8.25 5v3.25H5A1.75 1.75 0 0 0 3.25 10v4c0 .966.784 1.75 1.75 1.75h3.25V19c0 .966.784 1.75 1.75 1.75h4A1.75 1.75 0 0 0 15.75 19v-3.25H19A1.75 1.75 0 0 0 20.75 14v-4A1.75 1.75 0 0 0 19 8.25h-3.25V5A1.75 1.75 0 0 0 14 3.25h-4Z"/>`),
		g.Group(children),
	)
}