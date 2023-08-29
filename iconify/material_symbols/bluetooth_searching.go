package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothSearching(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 22v-7.6L4.4 19L3 17.6L8.6 12L3 6.4L4.4 5L9 9.6V2h1l5.7 5.7l-4.3 4.3l4.3 4.3L10 22H9Zm2-12.4l1.9-1.9L11 5.85V9.6Zm0 8.55l1.9-1.85l-1.9-1.9v3.75Zm5.55-3.8L14.25 12l2.3-2.3q.225.55.363 1.125T17.05 12q0 .6-.138 1.187t-.362 1.163Zm2.95 2.85L18.25 16q.5-.925.775-1.938T19.3 12q0-1.05-.275-2.063T18.25 8l1.25-1.25q.725 1.2 1.113 2.525T21 12q0 1.4-.388 2.713T19.5 17.2Z"/>`),
		g.Group(children),
	)
}