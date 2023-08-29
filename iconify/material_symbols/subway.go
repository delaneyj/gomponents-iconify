package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V8.85q0-2.125 1.1-3.688T6.2 2.8q1.35-.525 2.875-.662T12 2q1.4 0 2.925.138T17.8 2.8q2 .8 3.1 2.363T22 8.85V22H2Zm7.1-2l1.5-1.5h2.75l1.5 1.5h1.65v-.5l-1.05-1.05q1.1-.15 1.825-.988T18 15.5V9q0-1.95-1.75-2.475T12 6q-2.275 0-4.138.525T6 9v6.5q0 1.125.725 1.963t1.825.987L7.5 19.5v.5h1.6Zm-1.6-7V9h9v4h-9Zm8 3.5q-.425 0-.712-.288T14.5 15.5q0-.425.288-.713t.712-.287q.425 0 .713.288t.287.712q0 .425-.288.713t-.712.287Zm-7 0q-.425 0-.712-.288T7.5 15.5q0-.425.288-.713T8.5 14.5q.425 0 .713.288t.287.712q0 .425-.288.713T8.5 16.5Z"/>`),
		g.Group(children),
	)
}