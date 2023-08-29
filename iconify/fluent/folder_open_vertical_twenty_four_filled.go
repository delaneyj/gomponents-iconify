package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenVerticalTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 5.25A3.25 3.25 0 0 0 17.75 2H8.84l3.54 2.045a4.25 4.25 0 0 1 2.126 3.68V20h.745a3.25 3.25 0 0 0 3.25-3.25v-5.19l1.841-1.84a2.25 2.25 0 0 0 .66-1.591V5.25Zm-7.995 2.475v13.02c0 1.732-1.875 2.814-3.375 1.948l-5.26-3.036a2.75 2.75 0 0 1-1.375-2.382V4.255c0-1.731 1.875-2.814 3.375-1.948l5.259 3.037a2.75 2.75 0 0 1 1.375 2.381Z"/>`),
		g.Group(children),
	)
}