package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Node(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12" cy="3" r="1" fill="currentColor"/><circle cx="20" cy="8" r="1" fill="currentColor"/><circle cx="20" cy="16" r="1" fill="currentColor"/><circle cx="4" cy="8" r="1" fill="currentColor"/><circle cx="4" cy="16" r="1" fill="currentColor"/><path fill="currentColor" d="M20 14v-4a1.992 1.992 0 0 1-1.481-3.333l-4.783-2.69a1.983 1.983 0 0 1-3.472 0l-4.783 2.69A1.992 1.992 0 0 1 4 10v4a1.992 1.992 0 0 1 1.481 3.333l4.783 2.69a1.991 1.991 0 0 1 1.236-.952v-5.142a2 2 0 1 1 1 0v5.142a1.991 1.991 0 0 1 1.236.953l4.783-2.69A1.992 1.992 0 0 1 20 14Z"/><circle cx="12" cy="21" r="1" fill="currentColor"/><circle cx="12" cy="12" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}