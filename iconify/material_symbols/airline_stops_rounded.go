package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineStopsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 17q-.35-3.175-2.563-5.363T3.1 9.075q-.45-.05-.775-.35T2 8q0-.425.3-.712t.725-.238q2.925.275 5.313 1.925T12 13.3q.95-2.025 2.5-3.588T17.975 7H15q-.425 0-.713-.288T14 6q0-.425.288-.713T15 5h5q.425 0 .713.288T21 6v5q0 .425-.288.713T20 12q-.425 0-.713-.288T19 11V8.7q-2.325 1.425-4 3.525T13 17h1q.425 0 .713.288T15 18q0 .425-.288.713T14 19h-4q-.425 0-.713-.288T9 18q0-.425.288-.713T10 17h1Z"/>`),
		g.Group(children),
	)
}