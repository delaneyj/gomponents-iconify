package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoDeleteOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 6v13V6Zm4.25 15H7q-.825 0-1.413-.588T5 19V6h-.025q-.425 0-.7-.288T4 5q0-.425.288-.713T5 4h4q0-.425.288-.713T10 3h4q.425 0 .713.288T15 4h4q.425 0 .713.288T20 5q0 .425-.288.713T19 6v4.3q-.425-.125-.988-.213T17 10V6H7v13h3.3q.15.525.4 1.038t.55.962ZM10 8q-.425 0-.713.288T9 9v7q0 .425.288.713T10 17q0-1.575.5-2.588L11 13.4V9q0-.425-.288-.713T10 8Zm3 3.25q.425-.275.963-.55T15 10.3V9q0-.425-.288-.713T14 8q-.425 0-.713.288T13 9v2.25ZM17 22q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q1.025 0 1.938.4t1.587 1.075q.675.675 1.075 1.588T22 17q0 2.075-1.463 3.538T17 22Zm.5-5.2v-2.3q0-.2-.15-.35T17 14q-.2 0-.35.15t-.15.35v2.275q0 .2.075.388t.225.337l1.5 1.5q.15.15.35.15T19 19q.15-.15.15-.35T19 18.3l-1.5-1.5Z"/>`),
		g.Group(children),
	)
}