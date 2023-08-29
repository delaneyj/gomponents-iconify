package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20H6q-.425 0-.713-.288T5 19v-7H3.3q-.35 0-.475-.325t.15-.55l8.35-7.525q.275-.275.675-.275t.675.275L16 6.6V4.5q0-.2.15-.35T16.5 4h2q.2 0 .35.15t.15.35v4.8l2.025 1.825q.275.225.15.55T20.7 12H19v7q0 .425-.288.713T18 20h-4v-5q0-.425-.288-.713T13 14h-2q-.425 0-.713.288T10 15v5Zm0-9.975h4q0-.8-.6-1.313T12 8.2q-.8 0-1.4.513t-.6 1.312Z"/>`),
		g.Group(children),
	)
}