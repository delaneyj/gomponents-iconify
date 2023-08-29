package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditUndoLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 14.25V19h4.75l8.33-8.33l-5.27-4.23zM13 2.86V0L8 4l5 4V5h.86c2.29 0 4 1.43 4 4.29H20a6.51 6.51 0 0 0-6.14-6.43z"/>`),
		g.Group(children),
	)
}