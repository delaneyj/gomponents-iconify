package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontSize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.782 2.217a.4.4 0 0 0-.565 0l-2 2a.4.4 0 0 0 .565.566L2.1 3.466v8.068L.782 10.217a.4.4 0 0 0-.565.566l2 2a.4.4 0 0 0 .565 0l2-2a.4.4 0 0 0-.565-.566l-1.318 1.317V3.466l1.318 1.317a.4.4 0 0 0 .565-.566l-2-2Zm7.718.533a.5.5 0 0 1 .47.33l3 8.32a.5.5 0 0 1-.94.34l-.982-2.724H8.952l-.982 2.723a.5.5 0 0 1-.94-.34l3-8.319a.5.5 0 0 1 .47-.33Zm0 1.974l1.241 3.442H9.26l1.24-3.442Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}