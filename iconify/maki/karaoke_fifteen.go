package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KaraokeFifteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M12.1 2.952a2.988 2.988 0 0 0-5.11 1.965l3.142 3.142A2.988 2.988 0 0 0 12.1 2.952z" fill="currentColor"/><path d="M4.672 8.255L2.55 10.377a1 1 0 0 0 0 1.414l.707.707a1 1 0 0 0 1.414 0l2.121-2.121l2.122-2.122l-2.121-2.121zm.741 2.087l-.707-.707l2.087-2.087l.707.707z" fill="currentColor"/>`),
		g.Group(children),
	)
}