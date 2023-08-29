package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 1.28A1 1 0 0 0 8.35 1H2a1 1 0 0 0-1 1v6.35a1 1 0 0 0 .28.65L11 18.72a1 1 0 0 0 1.37 0l6.38-6.38a1 1 0 0 0-.03-1.34zM5 7a2 2 0 1 1 2-2a2 2 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}