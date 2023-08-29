package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PackageTwoOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19.425v-6.85L5 9.1v6.85l6 3.475Zm2 0l6-3.475V9.1l-6 3.475v6.85ZM12 22.3l-9-5.175V6.875L12 1.7l9 5.175v10.25L12 22.3Zm4-13.775l1.925-1.1L12 4l-1.95 1.125l5.95 3.4Zm-4 2.325l1.95-1.125L8.025 6.3l-1.95 1.125L12 10.85Z"/>`),
		g.Group(children),
	)
}