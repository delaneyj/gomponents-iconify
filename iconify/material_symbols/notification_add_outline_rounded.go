package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationAddOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22Zm-7-3q-.425 0-.713-.288T4 18q0-.425.288-.713T5 17h1v-7q0-2.075 1.25-3.688T10.5 4.2v-.7q0-.625.438-1.063T12 2q.625 0 1.063.438T13.5 3.5v.7q.45.125.875.288t.8.412q-.325.4-.563.837t-.437.913q-.475-.3-1.025-.475T12 6q-1.65 0-2.825 1.175T8 10v7h8v-3.075q.425.375.925.688t1.075.512V17h1q.425 0 .713.288T20 18q0 .425-.288.713T19 19H5Zm7-7.5Zm7-1.5h-2q-.425 0-.713-.288T16 9q0-.425.288-.713T17 8h2V6q0-.425.288-.713T20 5q.425 0 .713.288T21 6v2h2q.425 0 .713.288T24 9q0 .425-.288.713T23 10h-2v2q0 .425-.288.713T20 13q-.425 0-.713-.288T19 12v-2Z"/>`),
		g.Group(children),
	)
}