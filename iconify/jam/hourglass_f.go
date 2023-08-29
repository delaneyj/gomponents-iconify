package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 18a1 1 0 0 1 0 2H1a1 1 0 0 1 0-2v-4a4.99 4.99 0 0 1 2-4a4.992 4.992 0 0 1-2-4V2a1 1 0 1 1 0-2h10a1 1 0 0 1 0 2v4a4.992 4.992 0 0 1-2 4a4.99 4.99 0 0 1 2 4v4z"/>`),
		g.Group(children),
	)
}