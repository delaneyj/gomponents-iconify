package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClimateMiniSplitRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 12q-.825 0-1.413-.588T2 10V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v6q0 .825-.588 1.413T20 12h-2V8q0-.825-.588-1.413T16 6H8q-.825 0-1.413.588T6 8v4H4Zm4 0V8h8v4H8Zm-.025 2q.375 0 .663.225t.287.575q0 1.725-1.163 2.962T4.875 19q-.375 0-.625-.288T4 18.025q0-.4.238-.7t.612-.3q.925 0 1.55-.662t.625-1.588q0-.35.287-.563T7.975 14Zm7.975 0q-.375 0-.663.225T15 14.8q0 1.725 1.163 2.963T19.05 19q.375 0 .625-.288t.25-.687q0-.4-.238-.7t-.612-.3q-.925 0-1.55-.662t-.625-1.588q0-.35-.287-.563T15.95 14ZM12 20q-.425 0-.713-.288T11 19v-4q0-.425.288-.713T12 14q.425 0 .713.288T13 15v4q0 .425-.288.713T12 20Z"/>`),
		g.Group(children),
	)
}