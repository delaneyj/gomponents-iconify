package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unsubscribe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 13l8-5V6l-8 5l-8-5v2l8 5Zm7 10q-2.075 0-3.538-1.463T14 18q0-2.075 1.463-3.538T19 13q2.075 0 3.538 1.463T24 18q0 2.075-1.463 3.538T19 23Zm-3-4.5h6v-1h-6v1ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v5.7q-.7-.35-1.45-.525T19 11q-2.925 0-4.963 2.037T12 18q0 .5.075 1t.225 1H4Z"/>`),
		g.Group(children),
	)
}