package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M15.567 16.946a4 4 0 1 1 6.867 4.107a4 4 0 0 1-6.867-4.107Zm0 0A24.001 24.001 0 0 0 9.06 13.77l-.41-.129m6.917-6.586a4 4 0 1 1 6.867-4.107a4 4 0 0 1-6.867 4.107Zm0 0A24 24 0 0 1 9.06 10.23l-.41.129m0 0c.225.5.35 1.055.35 1.64s-.125 1.14-.35 1.64m0-3.28a4 4 0 1 0 0 3.28"/>`),
		g.Group(children),
	)
}