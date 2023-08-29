package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17a3 3 0 1 0 6 0a3 3 0 0 0-6 0m6 0V4h10v9.5M9 8h10m-4 11l2 2l4-4"/>`),
		g.Group(children),
	)
}