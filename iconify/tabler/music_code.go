package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17a3 3 0 1 0 6 0a3 3 0 0 0-6 0m6 0V4h10v9M9 8h10m1 13l2-2l-2-2m-3 0l-2 2l2 2"/>`),
		g.Group(children),
	)
}