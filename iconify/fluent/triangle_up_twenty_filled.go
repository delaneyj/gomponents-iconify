package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleUpTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.32 2.785c-.568-1.047-2.071-1.047-2.638 0l-6.5 12a1.5 1.5 0 0 0 1.32 2.213H16.5a1.5 1.5 0 0 0 1.319-2.214l-6.5-11.999Z"/>`),
		g.Group(children),
	)
}