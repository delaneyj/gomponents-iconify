package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationPause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.02 20.04L10 14l-7-3.5a.55.55 0 0 1 0-1L21 3l-3.634 10.062M17 17v5m4-5v5"/>`),
		g.Group(children),
	)
}