package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangularFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 2a2 2 0 0 0-2 2v26h4V4a2 2 0 0 0-2-2Zm21.626 11.056L9 20.596V4.404l18.626 7.539a.6.6 0 0 1 0 1.112Z"/>`),
		g.Group(children),
	)
}