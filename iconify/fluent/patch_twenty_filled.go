package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatchTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m8.267 3.06l1.026 1.025l-5.208 5.208L3.06 8.267a2.5 2.5 0 0 1 0-3.535L4.732 3.06a2.5 2.5 0 0 1 3.535 0Zm3.465 13.879l-1.025-1.025l5.207-5.207l1.025 1.025a2.5 2.5 0 0 1 0 3.536l-1.671 1.671a2.5 2.5 0 0 1-3.536 0ZM3.06 11.732a2.5 2.5 0 0 0 0 3.536l1.671 1.671a2.5 2.5 0 0 0 3.536 0l8.671-8.672a2.5 2.5 0 0 0 0-3.535l-1.67-1.672a2.5 2.5 0 0 0-3.536 0l-8.671 8.672ZM10 8.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1ZM8.5 10a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm3.5.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1ZM10.5 12a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/>`),
		g.Group(children),
	)
}