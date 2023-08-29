package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveUpRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 20q-2.925 0-4.963-2.038T1 13q0-2.675 1.763-4.663T7.175 6.05l-.9-.95Q6 4.8 6 4.4t.275-.675q.3-.3.725-.3t.7.275l2.6 2.6q.3.3.3.7t-.3.7l-2.575 2.575q-.3.3-.712.3t-.713-.3q-.275-.3-.275-.713t.275-.687l.775-.775Q5.3 8.45 4.15 9.825T3 13q0 2.075 1.463 3.538T8 18h2q.425 0 .713.288T11 19q0 .425-.288.713T10 20H8Zm6-9q-.425 0-.713-.288T13 10V5q0-.425.288-.713T14 4h7q.425 0 .713.288T22 5v5q0 .425-.288.713T21 11h-7Zm0 9q-.425 0-.713-.288T13 19v-5q0-.425.288-.713T14 13h7q.425 0 .713.288T22 14v5q0 .425-.288.713T21 20h-7Zm1-2h5v-3h-5v3Z"/>`),
		g.Group(children),
	)
}