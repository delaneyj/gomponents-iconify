package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneEventBusy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5h14v2H5z" opacity=".3"/><path fill="currentColor" d="M19 3h-1V1h-2v2H8V1H6v2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 16H5V9h14v10zm0-12H5V5h14v2zM9.29 17.47l2.44-2.44l2.44 2.44l1.06-1.06l-2.44-2.44l2.44-2.44l-1.06-1.06l-2.44 2.44l-2.44-2.44l-1.06 1.06l2.44 2.44l-2.44 2.44z"/>`),
		g.Group(children),
	)
}