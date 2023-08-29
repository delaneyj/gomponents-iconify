package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.11 21.46l-1.27 1.27L19.11 21H4a2 2 0 0 1-2-2V8c0-1.11.89-2 2-2h.11l-3-3l1.28-1.27l19.72 19.73M22 18.8L8 4.8V4c0-1.11.89-2 2-2h4a2 2 0 0 1 2 2v2h4a2 2 0 0 1 2 2v10.8M14 4h-4v2h4V4Z"/>`),
		g.Group(children),
	)
}