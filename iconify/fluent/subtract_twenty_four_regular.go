package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtractTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.754 12.5h16.492a.75.75 0 0 0 0-1.5H3.754a.75.75 0 0 0 0 1.5Z"/>`),
		g.Group(children),
	)
}