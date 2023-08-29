package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsInputComponentOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 23q-.425 0-.713-.288T3 22v-3.2q-.875-.3-1.438-1.063T1 16V7q0-.425.288-.713T2 6h1V2q0-.425.288-.713T4 1q.425 0 .713.288T5 2v4h1q.425 0 .713.288T7 7v9q0 .975-.563 1.738T5 18.8v3.225q0 .425-.288.7T4 23Zm8 0q-.425 0-.713-.288T11 22v-3.2q-.875-.3-1.438-1.063T9 16V7q0-.425.288-.713T10 6h1V2q0-.425.288-.713T12 1q.425 0 .713.288T13 2v4h1q.425 0 .713.288T15 7v9q0 .975-.563 1.738T13 18.8v3.225q0 .425-.288.7T12 23Zm8 0q-.425 0-.713-.288T19 22v-3.2q-.875-.3-1.438-1.063T17 16V7q0-.425.288-.713T18 6h1V2q0-.425.288-.713T20 1q.425 0 .713.288T21 2v4h1q.425 0 .713.288T23 7v9q0 .975-.563 1.738T21 18.8v3.225q0 .425-.288.7T20 23ZM3 8v4h2V8H3Zm8 0v4h2V8h-2Zm8 0v4h2V8h-2ZM4 17q.425 0 .713-.288T5 16v-2H3v2q0 .425.288.713T4 17Zm8 0q.425 0 .713-.288T13 16v-2h-2v2q0 .425.288.713T12 17Zm8 0q.425 0 .713-.288T21 16v-2h-2v2q0 .425.288.713T20 17ZM4 13Zm8 0Zm8 0ZM3 12h2h-2Zm8 0h2h-2Zm8 0h2h-2ZM4 14H3h2h-1Zm8 0h-1h2h-1Zm8 0h-1h2h-1Z"/>`),
		g.Group(children),
	)
}