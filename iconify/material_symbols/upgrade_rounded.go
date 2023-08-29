package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpgradeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 20q-.425 0-.713-.288T7 19q0-.425.288-.713T8 18h8q.425 0 .713.288T17 19q0 .425-.288.713T16 20H8Zm4-4q-.425 0-.713-.288T11 15V7.825L9.1 9.7q-.275.275-.688.288T7.7 9.7q-.275-.275-.275-.7t.275-.7l3.6-3.6q.15-.15.325-.212T12 4.425q.2 0 .375.063t.325.212l3.6 3.6q.275.275.288.688T16.3 9.7q-.275.275-.7.275t-.7-.275L13 7.825V15q0 .425-.287.713T12 16Z"/>`),
		g.Group(children),
	)
}