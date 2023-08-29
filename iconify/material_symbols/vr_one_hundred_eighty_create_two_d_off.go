package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrOneHundredEightyCreateTwoDOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17.75q-2.675-.675-4.338-2.838T2 10q0-1.125.3-2.15t.825-1.925l-2.5-2.5l1.4-1.4L21.2 21.2l-1.4 1.4l-.6-.6H12q-.825 0-1.413-.588T10 20v-7.2l-1.875-1.85q-.075.275-.1.525T8 12v5.75ZM10 2q2.75 0 4.913 1.663T17.75 8H12q-.275 0-.525.025t-.525.1l-5-5.025q.925-.55 1.95-.825T10 2Zm2.825 8H20q.825 0 1.413.588T22 12v7.175L12.825 10Zm3.375 9l-.7-.725l-1.425-1.4L12.5 19h3.7Z"/>`),
		g.Group(children),
	)
}