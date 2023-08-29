package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blinds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21v-2h2V3h16v16h2v2H2ZM6 7h8V5H6v2Zm0 4h8V9H6v2Zm0 8h12v-6h-2v1.825q.35.25.55.625t.2.8q0 .725-.513 1.238T15 18q-.725 0-1.238-.513t-.512-1.237q0-.425.2-.8t.55-.625V13H6v6ZM16 7h2V5h-2v2Zm0 4h2V9h-2v2Z"/>`),
		g.Group(children),
	)
}