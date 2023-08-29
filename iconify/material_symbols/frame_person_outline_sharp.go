package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FramePersonOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22v-6h2v4h4v2H2ZM2 8V2h6v2H4v4H2Zm14 14v-2h4v-4h2v6h-6Zm4-14V4h-4V2h6v6h-2Zm-8 4q-1.275 0-2.138-.863T9 9q0-1.25.863-2.125T12 6q1.25 0 2.125.875T15 9q0 1.275-.875 2.138T12 12Zm0-2q.425 0 .713-.288T13 9q0-.425-.288-.713T12 8q-.425 0-.713.288T11 9q0 .425.288.713T12 10Zm-6 8v-1.9q0-.525.263-.988t.712-.737q1.15-.675 2.413-1.025T12 13q1.35 0 2.613.35t2.412 1.025q.45.275.713.738T18 16.1V18H6Zm6-3q-1.025 0-2 .25T8.15 16h7.7q-.875-.5-1.85-.75T12 15Zm0-6Zm0 7h3.85h-7.7H12Z"/>`),
		g.Group(children),
	)
}