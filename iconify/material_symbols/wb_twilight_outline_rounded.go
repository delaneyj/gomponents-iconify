package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WbTwilightOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.8 7.25q.275.275.275.7t-.275.7l-.725.725q-.3.3-.713.3t-.712-.3q-.3-.3-.287-.713t.312-.712l.725-.725q.3-.275.713-.263t.687.288ZM3 20q-.425 0-.713-.288T2 19q0-.425.288-.713T3 18h18q.425 0 .713.288T22 19q0 .425-.288.713T21 20H3Zm9-16q.425 0 .713.288T13 5v1q0 .425-.288.713T12 7q-.425 0-.713-.288T11 6V5q0-.425.288-.713T12 4ZM4.25 7.2q.275-.275.7-.275t.7.275l.725.725q.3.3.3.713t-.3.712q-.3.275-.725.275t-.7-.3L4.225 8.6q-.275-.3-.263-.713T4.25 7.2ZM7.425 14h9.15q-.575-1.35-1.8-2.175T12 11q-1.55 0-2.775.825T7.425 14ZM5 16q0-2.925 2.038-4.963T12 9q2.925 0 4.963 2.038T19 16H5Zm7-2Z"/>`),
		g.Group(children),
	)
}