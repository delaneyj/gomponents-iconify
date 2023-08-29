package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationInstance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 6v14a2 2 0 0 0 2 2h18a2.006 2.006 0 0 0 2-2.007V6Zm8 13l.003-4.502L9 10l7 4.5ZM1.01 3.007L1 5h22l-.01-1.993A2 2 0 0 0 21 1H3a2 2 0 0 0-1.99 2.007Zm2.997 1a1 1 0 1 1 .999-.999a1 1 0 0 1-1 1Zm2.997 0a1 1 0 1 1 1-1a1.002 1.002 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}