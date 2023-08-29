package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalGasStationSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21V3h10v9h3v7.5h2v-8.2q-.225.125-.475.163T18 11.5q-1.05 0-1.775-.725T15.5 9q0-.8.438-1.438T17.1 6.65L15 4.55l1.05-1.05l3.7 3.6q.375.375.563.875T20.5 9v12h-5v-7.5H14V21H4Zm2-11h6V5H6v5Zm12 0q.425 0 .713-.288T19 9q0-.425-.288-.713T18 8q-.425 0-.713.288T17 9q0 .425.288.713T18 10Z"/>`),
		g.Group(children),
	)
}