package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanPersonTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6a3 3 0 0 1 3-3h1.5a.5.5 0 0 1 0 1H6a2 2 0 0 0-2 2v1.5a.5.5 0 0 1-1 0V6Zm9-2.5a.5.5 0 0 1 .5-.5H14a3 3 0 0 1 3 3v1.5a.5.5 0 0 1-1 0V6a2 2 0 0 0-2-2h-1.5a.5.5 0 0 1-.5-.5ZM3.5 12a.5.5 0 0 1 .5.5V14a2 2 0 0 0 1.098 1.785A1.987 1.987 0 0 0 6 16h1.5a.5.5 0 0 1 0 1H6a3 3 0 0 1-3-3v-1.5a.5.5 0 0 1 .5-.5Zm13 0a.5.5 0 0 1 .5.5V14a3 3 0 0 1-3 3h-1.5a.5.5 0 0 1 0-1H14a2 2 0 0 0 2-2v-1.5a.5.5 0 0 1 .5-.5ZM5.025 14.224A1 1 0 0 0 6 15v-.5a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 .5.5v.5a1 1 0 0 0 .975-.776A1.5 1.5 0 0 0 13.5 13h-7a1.5 1.5 0 0 0-1.475 1.224ZM12.75 9.25a2.75 2.75 0 1 0-5.5 0a2.75 2.75 0 0 0 5.5 0Zm-4.5 0a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0Z"/>`),
		g.Group(children),
	)
}