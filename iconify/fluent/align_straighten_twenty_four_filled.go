package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignStraightenTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.55 5.164A3.25 3.25 0 0 0 3.08 9.04l.268 1.21h16.534l-1.043-4.704a3.25 3.25 0 0 0-3.876-2.47L5.549 5.165Zm-.383 13.29L4.124 13.75h16.534l.268 1.21a3.25 3.25 0 0 1-2.47 3.876l-9.413 2.087a3.25 3.25 0 0 1-3.876-2.47ZM1 12a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5A.75.75 0 0 1 1 12Zm6 0a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5A.75.75 0 0 1 7 12Zm6.75-.75a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5h-2.5ZM19 12a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5A.75.75 0 0 1 19 12Z"/>`),
		g.Group(children),
	)
}