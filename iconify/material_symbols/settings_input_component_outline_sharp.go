package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsInputComponentOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23v-4.2H1V6h2V2q0-.425.288-.713T4 1q.425 0 .713.288T5 2v4h2v12.8H5V23H3Zm8 0v-4.2H9V6h2V2q0-.425.288-.713T12 1q.425 0 .713.288T13 2v4h2v12.8h-2V23h-2Zm8 0v-4.2h-2V6h2V2q0-.425.288-.713T20 1q.425 0 .713.288T21 2v4h2v12.8h-2V23h-2ZM3 8v4h2V8H3Zm8 0v4h2V8h-2Zm8 0v4h2V8h-2ZM3 17h2v-3H3v3Zm8 0h2v-3h-2v3Zm8 0h2v-3h-2v3ZM4 13Zm8 0Zm8 0ZM3 12h2h-2Zm8 0h2h-2Zm8 0h2h-2ZM3 14h2h-2Zm8 0h2h-2Zm8 0h2h-2Z"/>`),
		g.Group(children),
	)
}