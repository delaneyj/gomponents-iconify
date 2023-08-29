package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 2H7a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2ZM7 3a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm0 18a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm5-7a4.973 4.973 0 0 1-2-.422V11H7.422A4.997 4.997 0 1 1 12 14Zm5 7a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm0-16a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/><circle cx="12" cy="9" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}