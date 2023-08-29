package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WritingSignOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 19c3.333-2 5-4 5-6c0-3-1-3-2-3s-2.032 1.085-2 3c.034 2.048 1.658 2.877 2.5 4C8 19 9 19.5 10 18c.667-1 1.167-1.833 1.5-2.5c1 2.333 2.333 3.5 4 3.5H18m-2-3v1l2 2l.5-.5M20 16V5c0-1.121-.879-2-2-2s-2 .879-2 2v7m0-5h4M3 3l18 18"/>`),
		g.Group(children),
	)
}