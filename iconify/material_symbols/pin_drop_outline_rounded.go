package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinDropOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16.475q2.475-2 3.738-3.85T17 9.15q0-1.4-.513-2.388t-1.262-1.6q-.75-.612-1.625-.887T12 4q-.725 0-1.6.275t-1.625.888q-.75.612-1.262 1.6T7 9.15q0 1.625 1.263 3.475T12 16.475Zm0 2.125q-.25 0-.488-.075t-.437-.225q-2.025-1.575-4.05-3.963T5 9.15q0-1.775.638-3.113T7.274 3.8q1-.9 2.25-1.35T12 2q1.225 0 2.475.45t2.25 1.35q1 .9 1.638 2.238T19 9.15q0 2.8-2.025 5.188t-4.05 3.962q-.2.15-.438.225T12 18.6Zm0-7.6q.825 0 1.413-.588T14 9q0-.825-.588-1.413T12 7q-.825 0-1.413.588T10 9q0 .825.588 1.413T12 11ZM6 22q-.425 0-.713-.288T5 21q0-.425.288-.713T6 20h12q.425 0 .713.288T19 21q0 .425-.288.713T18 22H6Zm6-12.85Z"/>`),
		g.Group(children),
	)
}