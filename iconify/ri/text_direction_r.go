package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDirectionR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 5v10H9v-4a4 4 0 1 1 0-8h8v2h-2v10h-2V5h-2ZM9 5a2 2 0 1 0 0 4V5ZM7 17h12v2H7v2.5L3 18l4-3.5V17Z"/>`),
		g.Group(children),
	)
}