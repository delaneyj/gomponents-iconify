package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WritingOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7h4m-4 9v1l2 2l.5-.5M20 16V5c0-1.121-.879-2-2-2s-2 .879-2 2v7m2 7H5a2 2 0 1 1 0-4h4a2 2 0 1 0 0-4H6M3 3l18 18"/>`),
		g.Group(children),
	)
}