package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LatinCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#9266CC" d="M36 32a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28z"/><path fill="#FFF" d="M26 10h-6V4a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v6h-6a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h6v18a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V14h6a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1z"/>`),
		g.Group(children),
	)
}