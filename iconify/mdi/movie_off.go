package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.11 21.46l-1.27 1.27L18.11 20H4a2 2 0 0 1-2-2V6c0-.58.25-1.1.65-1.46L1.11 3l1.28-1.27l19.72 19.73m-.23-2.78c.08-.21.12-.44.12-.68V4h-4l2 4h-3l-2-4h-2l2 4h-3l-2-4H8l.8 1.6l13.08 13.08Z"/>`),
		g.Group(children),
	)
}