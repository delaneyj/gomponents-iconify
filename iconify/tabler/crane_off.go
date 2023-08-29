package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CraneOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 21h6m-3 0V9m0-4V3L8 4M6 6L3 9h6m4 0h8M9 3l10 6m-2 0v4a2 2 0 0 1 2 2m-2 2a2 2 0 0 1-2-2M3 3l18 18"/>`),
		g.Group(children),
	)
}