package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsCellSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 24q-.425 0-.713-.288T7 23q0-.425.288-.713T8 22q.425 0 .713.288T9 23q0 .425-.288.713T8 24Zm4 0q-.425 0-.713-.288T11 23q0-.425.288-.713T12 22q.425 0 .713.288T13 23q0 .425-.288.713T12 24Zm4 0q-.425 0-.713-.288T15 23q0-.425.288-.713T16 22q.425 0 .713.288T17 23q0 .425-.288.713T16 24ZM6 20V0h12v20H6Zm2-5h8V5H8v10Z"/>`),
		g.Group(children),
	)
}