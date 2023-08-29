package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrOneHundredEightyCreateTwoDOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17.75q-2.675-.675-4.338-2.838T2 10q0-1.125.3-2.15t.825-1.925l-1.8-1.8q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275L20.5 20.5q.3.3.263.738T20.4 22q-.25.25-.6.25t-.6-.25H12q-.825 0-1.413-.588T10 20v-7.2l-1.875-1.85q-.075.275-.1.525T8 12v5.75ZM10 2q2.75 0 4.913 1.663T17.75 8H12q-.275 0-.525.025t-.525.1l-5-5.025q.925-.55 1.95-.825T10 2Zm2.825 8H20q.825 0 1.413.588T22 12v7.175L12.825 10ZM13 19h3.2l-.7-.725l-1.425-1.4L12.8 18.6q-.1.125-.025.263T13 19Z"/>`),
		g.Group(children),
	)
}