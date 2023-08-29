package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsInputComponentSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 23v-4.2h-2V14h6v4.8h-2V23h-2Zm-8 0v-4.2H9V14h6v4.8h-2V23h-2Zm-8 0v-4.2H1V14h6v4.8H5V23H3ZM1 12V6h2V2q0-.425.288-.713T4 1q.425 0 .713.288T5 2v4h2v6H1Zm8 0V6h2V2q0-.425.288-.713T12 1q.425 0 .713.288T13 2v4h2v6H9Zm8 0V6h2V2q0-.425.288-.713T20 1q.425 0 .713.288T21 2v4h2v6h-6Z"/>`),
		g.Group(children),
	)
}