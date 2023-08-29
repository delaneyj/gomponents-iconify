package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlcoholShop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14 4h-4v3.5a2 2 0 0 0 1.5 1.93V13H11a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-.5V9.43A2 2 0 0 0 14 7.5V4zm-1 3.5a1 1 0 1 1-2 0V5h2v2.5zm-7.5-5V2a.5.5 0 0 0 0-1V.5A.5.5 0 0 0 5 0H4a.5.5 0 0 0-.5.5V1a.5.5 0 0 0 0 1v.5C3.5 3.93 1 5.57 1 7v6a1 1 0 0 0 1 1h5a1.1 1.1 0 0 0 1-1V7c0-1.35-2.5-3.15-2.5-4.5zm-1 9.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5z"/>`),
		g.Group(children),
	)
}