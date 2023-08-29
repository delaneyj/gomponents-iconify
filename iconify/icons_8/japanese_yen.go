package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseYen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m7.188 5l6.875 11H10v2h5v2h-5v2h5v5h2v-5h5v-2h-5v-2h5v-2h-4.063l6.875-11h-2.375L16 15.344L9.562 5H7.188z"/>`),
		g.Group(children),
	)
}