package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoctorSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 1.766a1.5 1.5 0 0 0-1.5 1.5v1.75H3.5a1.5 1.5 0 0 0-1.5 1.5v2.985a1.5 1.5 0 0 0 1.5 1.5H5v1.5a1.5 1.5 0 0 0 1.5 1.5h3a1.5 1.5 0 0 0 1.5-1.5v-1.5h1.5a1.5 1.5 0 0 0 1.5-1.5v-3a1.5 1.5 0 0 0-1.5-1.5H11V3.266a1.5 1.5 0 0 0-1.5-1.5h-3Z"/>`),
		g.Group(children),
	)
}