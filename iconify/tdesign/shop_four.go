package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShopFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 4a2 2 0 0 0-2 2v1h3V4H7Zm3 0v3h4V4h-4Zm6 0v3h3V6a2 2 0 0 0-2-2h-1Zm5 3h1v2h-1v11h1v2H2v-2h1V9H2V7h1V6a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4v1Zm-2 2H5v11h3v-5a4 4 0 0 1 8 0v5h3V9Zm-5 11v-5a2 2 0 1 0-4 0v5h4Z"/>`),
		g.Group(children),
	)
}