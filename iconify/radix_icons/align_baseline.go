package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBaseline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 1.75a.5.5 0 0 1 .47.33l3 8.32a.5.5 0 1 1-.94.34l-.982-2.724H8.952L7.97 10.74a.5.5 0 0 1-.94-.339l3-8.32a.5.5 0 0 1 .47-.33Zm0 1.974l1.241 3.442H9.26l1.24-3.442ZM2.5 2.1c.22 0 .4.18.4.4v7.034l1.317-1.317a.4.4 0 0 1 .565.566l-2 2a.4.4 0 0 1-.565 0l-2-2a.4.4 0 1 1 .565-.566L2.1 9.534V2.5c0-.22.18-.4.4-.4ZM.1 13.5a.4.4 0 0 1 .4-.4h14a.4.4 0 1 1 0 .8H.5a.4.4 0 0 1-.4-.4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}