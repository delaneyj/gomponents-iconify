package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FenceElectric(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 9v2H7V9H5v2H3V9H1v12h2v-2h2v2h2v-2h2v2h2v-2h2v2h2v-2h2v2h2v-2h2v2h2V9h-2v2h-2V9h-2v2h-2V9h-2v2h-2V9H9m-6 4h2v4H3v-4m4 0h2v4H7v-4m4 0h2v4h-2v-4m4 0h2v4h-2v-4m4 0h2v4h-2v-4M7 4h4V2l6 3h-4v2L7 4Z"/>`),
		g.Group(children),
	)
}