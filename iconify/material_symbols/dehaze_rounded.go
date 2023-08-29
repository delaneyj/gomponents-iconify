package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DehazeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 7q-.425 0-.713-.288T3 6q0-.425.288-.713T4 5h16q.425 0 .713.288T21 6q0 .425-.288.713T20 7H4Zm0 12q-.425 0-.713-.288T3 18q0-.425.288-.713T4 17h16q.425 0 .713.288T21 18q0 .425-.288.713T20 19H4Zm0-6q-.425 0-.713-.288T3 12q0-.425.288-.713T4 11h16q.425 0 .713.288T21 12q0 .425-.288.713T20 13H4Z"/>`),
		g.Group(children),
	)
}