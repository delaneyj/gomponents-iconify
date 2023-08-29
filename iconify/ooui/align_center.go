package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 15h18v2H1zM1 3h18v2H1z"/><rect width="8" height="6" x="6" y="7" fill="currentColor" rx="1"/>`),
		g.Group(children),
	)
}