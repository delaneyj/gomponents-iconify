package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrAr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 10a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2h-3.868a2 2 0 0 1-1.715-.971l-1.56-2.6a1 1 0 0 0-1.714 0l-1.56 2.6A2 2 0 0 1 7.868 19H4a2 2 0 0 1-2-2v-7Zm1.813-3.219A4 4 0 0 1 7.14 5h9.718a4 4 0 0 1 3.328 1.781L21 8H3l.813-1.219Z"/>`),
		g.Group(children),
	)
}