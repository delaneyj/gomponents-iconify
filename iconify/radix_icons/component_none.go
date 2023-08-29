package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.854 1.49a.5.5 0 0 0-.708 0L1.49 7.146a.5.5 0 0 0 0 .708l2.474 2.474l-2.318 2.318a.5.5 0 0 0 .708.708l2.318-2.318l2.474 2.474a.5.5 0 0 0 .708 0l5.657-5.656a.5.5 0 0 0 0-.708l-2.475-2.474l2.318-2.318a.5.5 0 1 0-.708-.708l-2.318 2.318L7.854 1.49ZM9.62 4.672L7.5 2.55L2.55 7.5l2.122 2.121l4.95-4.95Zm-4.24 5.656L7.5 12.45l4.95-4.95l-2.121-2.121l-4.95 4.95Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}