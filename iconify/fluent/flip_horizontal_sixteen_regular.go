package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipHorizontalSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.925 12.763A.5.5 0 0 1 14.5 13h-5a.5.5 0 0 1-.5-.5v-10a.5.5 0 0 1 .947-.224l5 10a.5.5 0 0 1-.022.487ZM10 4.618V12h3.691L10 4.618ZM1.5 13a.5.5 0 0 1-.447-.724l5-10A.5.5 0 0 1 7 2.5v10a.5.5 0 0 1-.5.5h-5Z"/>`),
		g.Group(children),
	)
}