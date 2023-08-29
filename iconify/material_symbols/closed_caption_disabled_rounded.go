package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptionDisabledRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.575L17.175 20H5q-.825 0-1.413-.588T3 18V5.825L1.375 4.2q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.7t-.3.7q-.3.3-.712.3t-.713-.3ZM7 15h3q.425 0 .713-.288T11 14v-.625q0-.2-.15-.35t-.35-.15H10q-.2 0-.35.15t-.15.35v.125h-2v-3.175L6.375 9.2v.025Q6.2 9.35 6.1 9.55T6 10v4q0 .425.288.713T7 15Zm14 3.125l-3.35-3.35q.175-.125.263-.338T18 14v-.5q0-.2-.15-.35T17.5 13H17q-.2 0-.35.15t-.15.35h-.125L14.5 11.625V10.5h2v.125q0 .2.15.35t.35.15h.5q.2 0 .35-.15t.15-.35V10q0-.425-.288-.713T17 9h-3q-.425 0-.713.288T13 10v.125L6.875 4H19q.825 0 1.413.588T21 6v12.125Z"/>`),
		g.Group(children),
	)
}