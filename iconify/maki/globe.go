package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Globe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m11.981 9.481l-.721-.721a4.979 4.979 0 1 1-7.02-7.02l-.721-.721A5.991 5.991 0 0 0 7 11.475V13h-.5a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1H8v-1.527a5.983 5.983 0 0 0 3.981-1.992ZM7.5 9A3.5 3.5 0 1 0 4 5.5A3.5 3.5 0 0 0 7.5 9Zm1-5l.364-.592a2.5 2.5 0 0 1 .823.9L9.5 5h-1Zm-1.844-.844L7.5 4v1L8 6h1.94a2.494 2.494 0 0 1-1.5 1.814L8 7H6.5L5.05 5a2.5 2.5 0 0 1 1.606-1.844Z"/>`),
		g.Group(children),
	)
}