package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MopOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 11h2V4q0-.425-.288-.713T12 3q-.425 0-.713.288T11 4v7Zm-6 4h14v-2H5v2Zm-1.45 6H6v-3h2v3h3v-3h2v3h3v-3h2v3h2.45l-1-4H4.55l-1 4ZM1 23l2-8v-4h6V4q0-1.25.875-2.125T12 1q1.25 0 2.125.875T15 4v7h6v4l2 8H1Zm18-10H5h14Zm-6-2h-2h2Z"/>`),
		g.Group(children),
	)
}