package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditUndoRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 15.25V20h4.75l8.33-8.33l-5.27-4.23z"/><path fill="currentColor" d="M13 2.86V0l5 4l-5 4V5h-.86c-2.28 0-4 1.43-4 4.29H6a6.51 6.51 0 0 1 6.14-6.43z"/>`),
		g.Group(children),
	)
}