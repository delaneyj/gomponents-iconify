package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionMarkCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm10.707-3.707A.994.994 0 0 0 11.998 8a.994.994 0 0 0-.705.293a1 1 0 0 1-1.414-1.414a2.994 2.994 0 0 1 2.116-.88a2.994 2.994 0 0 1 2.126.88A2.994 2.994 0 0 1 13 11.829l.001.167a1 1 0 0 1-2 .008l-.004-1A1 1 0 0 1 11.998 10a.994.994 0 0 0 .71-.293A.994.994 0 0 0 13 9a.994.994 0 0 0-.293-.707ZM10.5 16a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5V16Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}