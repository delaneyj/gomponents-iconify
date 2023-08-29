package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mountains(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.85 17.47l-5-8a1 1 0 0 0-1.7 0l-1 1.63l-3.29-5.6a1 1 0 0 0-1.72 0l-7 12A1 1 0 0 0 3 19h18a1 1 0 0 0 .85-1.53ZM10.45 17H4.74L10 8l2.93 5Zm2.35 0l2.2-3.43l1-1.68L19.2 17Z"/>`),
		g.Group(children),
	)
}