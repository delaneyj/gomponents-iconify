package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PediatricsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 7V5h4V2h2v3h4v2H7Zm0 15V11q0-1.25.875-2.125T10 8h4q1.25 0 2.125.875T17 11v11H7Zm2-2h6v-9q0-.425-.288-.713T14 10h-4q-.425 0-.713.288T9 11v1h3v2H9v2h3v2H9v2Zm0 0V10v10Z"/>`),
		g.Group(children),
	)
}