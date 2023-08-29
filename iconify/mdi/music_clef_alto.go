package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicClefAlto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 4h2v16H5m10.46-7h-.63l-1-1l1-1h.63a3.5 3.5 0 1 0-3.5-3.5h2a1.5 1.5 0 1 1 1.5 1.5H14l-2 2h-1V4H9v16h2v-7h1l2 2h1.46a1.5 1.5 0 1 1-1.5 1.5h-2a3.5 3.5 0 1 0 3.5-3.5Z"/>`),
		g.Group(children),
	)
}