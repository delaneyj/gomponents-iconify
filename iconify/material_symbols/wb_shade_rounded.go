package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WbShadeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.2 19.7l-4.875-4.875q-.275-.275-.275-.625t.275-.625q.275-.275.625-.275t.625.275l4.875 4.875q.275.275.275.625t-.275.625q-.275.275-.625.275T19.2 19.7ZM15 20q-.425 0-.712-.287T14 19v-2l3 3h-2ZM5 20q-.425 0-.712-.287T4 19v-9h-.8q-.35 0-.475-.3t.125-.55l4.8-4.8Q7.7 4.3 8 4.2q.05 0 .35.15l4.8 4.8q.25.25.125.55t-.475.3H12v9q0 .425-.288.713T11 20H5Zm2-6h2v-4H7v4Z"/>`),
		g.Group(children),
	)
}