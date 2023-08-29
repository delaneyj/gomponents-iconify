package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nutrition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 18a4 4 0 0 1-4 4h-4a4 4 0 0 1-4-4v-2h12v2M4 3h10a2 2 0 0 1 2 2v9H8v5H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2m0 3v2h2V6H4m10 2V6H8v2h6M4 10v2h2v-2H4m4 0v2h6v-2H8m-4 4v2h2v-2H4Z"/>`),
		g.Group(children),
	)
}