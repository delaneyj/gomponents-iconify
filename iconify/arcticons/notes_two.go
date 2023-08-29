package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotesTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.437 5.54a3.842 3.842 0 0 1 2.994 1.227l2.864 2.904a3.991 3.991 0 0 1 0 5.578l-5.339 5.338l-8.482-8.521l5.339-5.339A4.12 4.12 0 0 1 35.437 5.5v.04Zm.52 15.008l-18.96 19.018L5.57 42.5l2.957-11.425l18.949-19.01"/>`),
		g.Group(children),
	)
}