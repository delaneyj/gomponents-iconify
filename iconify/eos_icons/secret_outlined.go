package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SecretOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.172 19l2 2l-1 1H2v-1a2.002 2.002 0 0 1 2-2h.172M5 17H4a4 4 0 0 0-4 4v3h8l-1-1l2-2l-4-4Zm14.828 2H20a2.002 2.002 0 0 1 2 2v1h-3.172l-1-1l2-2M19 17l-4 4l2 2l-1 1h8v-3a4 4 0 0 0-4-4Zm-7 1l-1 2v4h2v-4l-1-2zm5-13V2a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v3H4v1h16V5Zm-2-1H9V2h6Zm.5 3a2.488 2.488 0 0 0-1.989 1H10.49a2.436 2.436 0 1 0 .46 1h2.101a2.5 2.5 0 1 0 2.45-2Zm-7 3.5a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm7 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}