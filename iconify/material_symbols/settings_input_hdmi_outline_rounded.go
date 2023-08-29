package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsInputHdmiOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 19l-3-6V7.975q0-.425.288-.7T6 7V4q0-.825.588-1.413T8 2h8q.825 0 1.413.588T18 4v3q.425 0 .713.275t.287.7V13l-3 6v2q0 .425-.288.713T15 22H9q-.425 0-.713-.288T8 21v-2ZM8 7h2V5.475q0-.2.15-.338T10.5 5q.2 0 .35.15t.15.35V7h2V5.475q0-.2.15-.338T13.5 5q.2 0 .35.15t.15.35V7h2V4H8v3Zm2 13h4v-1.5l3-6V9H7v3.5l3 6V20Zm2-5.5Z"/>`),
		g.Group(children),
	)
}