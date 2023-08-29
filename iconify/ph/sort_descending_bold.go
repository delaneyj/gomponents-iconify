package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortDescendingBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M36 128a12 12 0 0 1 12-12h68a12 12 0 0 1 0 24H48a12 12 0 0 1-12-12Zm12-52h52a12 12 0 0 0 0-24H48a12 12 0 0 0 0 24Zm132 104H48a12 12 0 0 0 0 24h132a12 12 0 0 0 0-24Zm52.49-100.49l-40-40a12 12 0 0 0-17 0l-40 40a12 12 0 0 0 17 17L172 77v67a12 12 0 0 0 24 0V77l19.51 19.52a12 12 0 0 0 17-17Z"/>`),
		g.Group(children),
	)
}