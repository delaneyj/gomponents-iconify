package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 18l-2-4l-7-3.5a.55.55 0 0 1 0-1L21 3l-3.616 10.015M16 22l5-5m0 4.5V17h-4.5"/>`),
		g.Group(children),
	)
}