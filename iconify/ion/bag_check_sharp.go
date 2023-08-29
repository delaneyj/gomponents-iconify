package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BagCheckSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M460 160h-88v-12A116.13 116.13 0 0 0 258.89 32h-5.78A116.13 116.13 0 0 0 140 148v12H52a4 4 0 0 0-4 4v300a16 16 0 0 0 16 16h384a16 16 0 0 0 16-16V164a4 4 0 0 0-4-4Zm-280-11c0-41.84 33.41-76.56 75.25-77A76.08 76.08 0 0 1 332 148v12H180Zm50.81 252.12l-61.37-71.72l24.31-20.81L230 350.91l87.51-109.4l25 20Z"/>`),
		g.Group(children),
	)
}