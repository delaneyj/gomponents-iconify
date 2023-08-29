package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubwayOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V8.85q0-2.125 1.1-3.688T6.2 2.8q1.35-.525 2.875-.662T12 2q1.4 0 2.925.138T17.8 2.8q2 .8 3.1 2.363T22 8.85V22H2Zm7.1-2h5.75l-1.5-1.5H10.6L9.1 20Zm-1.6-7h9V9h-9v4Zm8 3.5q.425 0 .713-.288t.287-.712q0-.425-.288-.713T15.5 14.5q-.425 0-.713.288t-.287.712q0 .425.288.713t.712.287Zm-7 0q.425 0 .713-.288T9.5 15.5q0-.425-.288-.713T8.5 14.5q-.425 0-.713.288T7.5 15.5q0 .425.288.713t.712.287ZM4 20h3.5v-.5l1.05-1.05q-1.1-.15-1.825-.988T6 15.5V9q0-1.95 1.863-2.475T12 6q2.5 0 4.25.525T18 9v6.5q0 1.125-.725 1.963t-1.825.987l1.05 1.05v.5H20V8.85q0-1.5-.738-2.563T17.05 4.65q-1.1-.425-2.438-.537T12 4q-1.275 0-2.613.113T6.95 4.65q-1.475.575-2.212 1.638T4 8.85V20Zm0 0h16H4Z"/>`),
		g.Group(children),
	)
}