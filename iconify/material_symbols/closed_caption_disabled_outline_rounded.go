package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptionDisabledOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.875 4H19q.825 0 1.413.588T21 6v12.125l-2-2V6H8.875l-2-2Zm9.625 9.5q0-.2.15-.35T17 13h.5q.2 0 .35.15t.15.35v.5q0 .225-.088.438t-.262.337L16.375 13.5h.125ZM14 9h3q.425 0 .713.288T18 10v.625q0 .2-.15.35t-.35.15H17q-.2 0-.35-.15t-.15-.35V10.5h-2v1.125l-1.5-1.5V10q0-.425.288-.713T14 9Zm-.05 2.05ZM10.1 12.9ZM9.025 9l1.5 1.5H7.5v3h2v-.125q0-.2.15-.35t.35-.15h.5q.2 0 .35.15t.15.35V14q0 .425-.288.713T10 15H7q-.425 0-.713-.288T6 14v-4q0-.425.288-.713T7 9h2.025ZM4.2 4.175L6.025 6H5v12h10.175L1.375 4.2q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.7t-.3.7q-.3.3-.712.3t-.713-.3L17.175 20H5q-.825 0-1.413-.588T3 18V6q0-.625.338-1.113t.862-.712Z"/>`),
		g.Group(children),
	)
}