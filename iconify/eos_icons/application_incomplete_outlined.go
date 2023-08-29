package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationIncompleteOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3v2h22l-.01-1.993A2 2 0 0 0 21 1H3a2.001 2.001 0 0 0-2 2Zm3 1a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm3 0a1 1 0 1 1 1-1a1.004 1.004 0 0 1-1 1Zm5 6a4 4 0 1 0 4 4h-4Z"/><path fill="currentColor" d="M1 6v14a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V6Zm20 14H3V8h18Z"/>`),
		g.Group(children),
	)
}