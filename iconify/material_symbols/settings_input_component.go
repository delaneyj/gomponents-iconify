package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsInputComponent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 12V6h2V2q0-.425.288-.713T4 1q.425 0 .713.288T5 2v4h2v6H1Zm2 11v-4.2q-.875-.3-1.438-1.063T1 16v-2h6v2q0 .975-.563 1.738T5 18.8V23H3Zm6-11V6h2V2q0-.425.288-.713T12 1q.425 0 .713.288T13 2v4h2v6H9Zm2 11v-4.2q-.875-.3-1.438-1.063T9 16v-2h6v2q0 .975-.563 1.738T13 18.8V23h-2Zm6-11V6h2V2q0-.425.288-.713T20 1q.425 0 .713.288T21 2v4h2v6h-6Zm2 11v-4.2q-.875-.3-1.438-1.063T17 16v-2h6v2q0 .975-.563 1.738T21 18.8V23h-2Z"/>`),
		g.Group(children),
	)
}