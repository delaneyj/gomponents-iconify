package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CutAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.236 8a3 3 0 1 0-4.472-4a3 3 0 0 0 4.472 4Zm0 0L16 16m1-4h1m4 0h1M6.236 16a3 3 0 1 1-4.472 4a3 3 0 0 1 4.472-4Zm0 0L16 8"/>`),
		g.Group(children),
	)
}