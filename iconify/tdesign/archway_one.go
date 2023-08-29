package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchwayOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v5h-2.915l.385 3H22v2h-2.274l1.393 10.865l-1.984.254L17.71 12H6.258l-1.39 11.116l-1.984-.248L4.242 12H2v-2h2.492l.375-3H2V2Zm4.883 5l-.375 3H11V7H6.883ZM13 7v3h4.454l-.385-3H13ZM4 4v1h16V4H4Z"/>`),
		g.Group(children),
	)
}