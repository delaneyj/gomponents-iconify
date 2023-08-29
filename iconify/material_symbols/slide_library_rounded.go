package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideLibraryRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V8q0-.425.288-.713T3 7q.425 0 .713.288T4 8v10h13q.425 0 .713.288T18 19q0 .425-.288.713T17 20H4Zm4-4q-.825 0-1.413-.588T6 14V6q0-.825.588-1.413T8 4h7.4l-2 2H8v8h12V8.6l2-2V14q0 .825-.588 1.413T20 16H8Zm5.5-4q-.475 0-.913-.163t-.787-.487q-.125-.125-.075-.262t.175-.188q.275-.075.438-.35T12.5 10q0-.625.438-1.062T14 8.5q.625 0 1.063.438T15.5 10q0 .825-.588 1.413T13.5 12Zm3.275-3L15 7.225l4.1-4.075q.15-.15.35-.15t.35.15l1.05 1.05q.15.15.15.35t-.15.35L16.775 9Z"/>`),
		g.Group(children),
	)
}