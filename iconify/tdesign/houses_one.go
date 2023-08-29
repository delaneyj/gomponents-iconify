package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HousesOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h4v-7h8v7h4V4H4Zm10 16v-5h-4v5h4ZM8 6v3H6V6h2Zm5 0v3h-2V6h2Zm5 0v3h-2V6h2Z"/>`),
		g.Group(children),
	)
}