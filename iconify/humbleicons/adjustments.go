package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adjustments(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v10m0 6v-.5M17.5 4v5m0 11v-5.56M6.5 4v2m0 14v-8.44M6.5 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm5.5 8a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm5.5-5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/>`),
		g.Group(children),
	)
}