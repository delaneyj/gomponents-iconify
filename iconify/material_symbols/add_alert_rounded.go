package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddAlertRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19q-.425 0-.713-.288T4 18q0-.425.288-.713T5 17h1v-7q0-2.075 1.25-3.688T10.5 4.2v-.7q0-.625.438-1.063T12 2q.625 0 1.063.438T13.5 3.5v.7q2 .5 3.25 2.113T18 10v7h1q.425 0 .713.288T20 18q0 .425-.288.713T19 19H5Zm7 3q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22Zm-1-9v1q0 .425.288.713T12 15q.425 0 .713-.288T13 14v-1h1q.425 0 .713-.288T15 12q0-.425-.288-.713T14 11h-1v-1q0-.425-.288-.713T12 9q-.425 0-.713.288T11 10v1h-1q-.425 0-.713.288T9 12q0 .425.288.713T10 13h1Z"/>`),
		g.Group(children),
	)
}