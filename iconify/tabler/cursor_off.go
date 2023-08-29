package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 4a3 3 0 0 1 3 3v1m0 9a3 3 0 0 1-3 3m6-16a3 3 0 0 0-3 3v1m0 4v5a3 3 0 0 0 3 3M3 3l18 18"/>`),
		g.Group(children),
	)
}