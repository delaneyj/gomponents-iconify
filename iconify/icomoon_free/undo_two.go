package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UndoTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11.904 16C13.681 12.781 13.98 7.87 7 8.034V12L1 6l6-6v3.881c8.359-.218 9.29 7.378 4.904 12.119z"/>`),
		g.Group(children),
	)
}