package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewColumnRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19q-.425 0-.713-.288T3 18V6q0-.425.288-.713T4 5h3.325q.425 0 .713.288T8.325 6v12q0 .425-.287.713T7.325 19H4Zm6.325 0q-.425 0-.713-.288T9.325 18V6q0-.425.288-.713T10.325 5h3.325q.425 0 .713.288T14.65 6v12q0 .425-.288.713T13.65 19h-3.325Zm6.325 0q-.425 0-.713-.288T15.65 18V6q0-.425.288-.713T16.65 5h3.325q.425 0 .713.288t.287.712v12q0 .425-.288.713t-.712.287H16.65Z"/>`),
		g.Group(children),
	)
}