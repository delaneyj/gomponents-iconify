package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptionDisabledOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.025 9l1.5 1.5H7.5v3h2V13H11v1q0 .425-.288.713T10 15H7q-.425 0-.713-.288T6 14v-4q0-.425.288-.713T7 9h2.025Zm-2.15-5H19q.825 0 1.413.588T21 6v12.125l-2-2V6H8.875l-2-2ZM18 13v1q0 .225-.088.438t-.262.337L16.375 13.5h.125V13H18Zm-1.5-2v-.5h-2v1.125l-1.5-1.5V10q0-.425.288-.713T14 9h3q.425 0 .713.288T18 10v1h-1.5Zm-2.55.05ZM10.1 12.9ZM4.2 4.175L6.025 6H5v12h10.175L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425l-3.3-3.3H5q-.825 0-1.413-.588T3 18V6q0-.625.338-1.113t.862-.712Z"/>`),
		g.Group(children),
	)
}