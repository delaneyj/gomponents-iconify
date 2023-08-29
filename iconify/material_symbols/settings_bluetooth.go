package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsBluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 24q-.425 0-.713-.288T7 23q0-.425.288-.713T8 22q.425 0 .713.288T9 23q0 .425-.288.713T8 24Zm4 0q-.425 0-.713-.288T11 23q0-.425.288-.713T12 22q.425 0 .713.288T13 23q0 .425-.288.713T12 24Zm4 0q-.425 0-.713-.288T15 23q0-.425.288-.713T16 22q.425 0 .713.288T17 23q0 .425-.288.713T16 24Zm-5-4v-7.6L6.4 17L5 15.6l5.6-5.6L5 4.4L6.4 3L11 7.6V0h1l5.7 5.7l-4.3 4.3l4.3 4.3L12 20h-1Zm2-3.85l1.9-1.85l-1.9-1.9v3.75Zm0-8.55l1.9-1.9L13 3.85V7.6Z"/>`),
		g.Group(children),
	)
}