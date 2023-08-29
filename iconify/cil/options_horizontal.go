package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OptionsHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M17.25 12a3 3 0 1 0 3-3a3.004 3.004 0 0 0-3 3zm4.5 0a1.5 1.5 0 1 1-1.5-1.5a1.502 1.502 0 0 1 1.5 1.5z" fill="currentColor"/><path d="M6.75 12a3 3 0 1 0-3 3a3.004 3.004 0 0 0 3-3zm-4.5 0a1.5 1.5 0 1 1 1.5 1.5a1.502 1.502 0 0 1-1.5-1.5z" fill="currentColor"/><path d="M15 12a3 3 0 1 0-3 3a3.004 3.004 0 0 0 3-3zm-4.5 0a1.5 1.5 0 1 1 1.5 1.5a1.502 1.502 0 0 1-1.5-1.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}