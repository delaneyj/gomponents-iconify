package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TooltipAccount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h4l4 4l4-4h4a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m-8 2.3c1.5 0 2.7 1.2 2.7 2.7c0 1.5-1.2 2.7-2.7 2.7c-1.5 0-2.7-1.2-2.7-2.7c0-1.5 1.2-2.7 2.7-2.7M18 15H6v-.9c0-2 4-3.1 6-3.1s6 1.1 6 3.1v.9Z"/>`),
		g.Group(children),
	)
}