package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionMarkSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-14Zm9.707 4.292A.994.994 0 0 0 11.998 8a.994.994 0 0 0-.705.293a1 1 0 0 1-1.414-1.414a2.994 2.994 0 0 1 2.116-.88a2.994 2.994 0 0 1 2.126.88A2.994 2.994 0 0 1 13 11.829l.001.167a1 1 0 0 1-2 .008l-.004-1A1 1 0 0 1 11.998 10a.994.994 0 0 0 .71-.293A.994.994 0 0 0 13 9a.994.994 0 0 0-.293-.707ZM10.5 16a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5V16Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}