package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightGroupOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22v-2h8v2H3Zm14-1q-.825 0-1.413-.588T15 19h-3q-.425 0-.713-.275t-.287-.7q-.05-2.25 1.4-3.912T16 12.1V2h2v10.1q2.175.35 3.613 2.013T23 18.024q0 .425-.288.7T22 19h-3q0 .825-.575 1.413T17 21ZM6 19v-8H2q-.5 0-.8-.388t-.175-.862l1.875-7q.1-.325.363-.537T3.875 2h6.25q.35 0 .613.213t.362.537l1.875 7q.125.475-.175.863T12 11H8v8H6Zm7.1-2h7.8q-.3-1.35-1.4-2.175T17 14q-1.375 0-2.475.825T13.1 17ZM3.3 9h7.4L9.375 4H4.65L3.3 9ZM7 6.5Zm10 9Z"/>`),
		g.Group(children),
	)
}