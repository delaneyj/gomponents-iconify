package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonRaisedHandRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23q-.425 0-.713-.288T2 22q0-.425.288-.713T3 21h18q.425 0 .713.288T22 22q0 .425-.288.713T21 23H3Zm2-3q-.425 0-.713-.288T4 19v-5q-.825-1.35-1.275-2.863t-.45-3.087q0-1.525.388-3t.912-2.9q.2-.525.65-.838t1-.312Q6 1 6.55 1.525T7 2.775L6.725 5.05q-.15 1.2.213 2.275t1.087 1.887q.725.813 1.75 1.3T12 11q1.5 0 3.013.313t2.637.887q1.125.575 1.738 1.463T20 15.85V19q0 .425-.288.713T19 20h-9v-.925q0-.85.575-1.463T12 17h3q.425 0 .713-.288T16 16q0-.425-.288-.713T15 15h-3q-1.675 0-2.838 1.2T8 19.075V20H5Zm7-10q-1.65 0-2.825-1.175T8 6q0-1.65 1.175-2.825T12 2q1.65 0 2.825 1.175T16 6q0 1.65-1.175 2.825T12 10Z"/>`),
		g.Group(children),
	)
}