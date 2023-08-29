package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func W(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.671 10.027a2 2 0 0 0-1.644 2.302l4 24a2 2 0 0 0 3.71.663L24 26.031l6.263 10.961a2 2 0 0 0 3.71-.663l4-24a2 2 0 1 0-3.946-.658l-3.077 18.46l-5.214-9.123a2 2 0 0 0-3.473 0l-5.214 9.124l-3.076-18.461a2 2 0 0 0-2.302-1.644Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}