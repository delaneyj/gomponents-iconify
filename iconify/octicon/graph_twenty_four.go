package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2.5 2.75a.75.75 0 0 0-1.5 0v18.5c0 .414.336.75.75.75H20a.75.75 0 0 0 0-1.5H2.5V2.75Z"/><path fill="currentColor" d="M22.28 7.78a.75.75 0 0 0-1.06-1.06l-5.72 5.72l-3.72-3.72a.75.75 0 0 0-1.06 0l-6 6a.75.75 0 1 0 1.06 1.06l5.47-5.47l3.72 3.72a.75.75 0 0 0 1.06 0l6.25-6.25Z"/>`),
		g.Group(children),
	)
}