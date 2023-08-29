package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.555 8.168l-9-6a1 1 0 0 0-1.109 0l-9 6A1 1 0 0 0 2 9v11c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V9c0-.334-.167-.646-.445-.832zM12 4.202L19.197 9L12 13.798L4.803 9L12 4.202zM4 20v-9.131l7.445 4.963a1 1 0 0 0 1.11 0L20 10.869L19.997 20H4z"/>`),
		g.Group(children),
	)
}