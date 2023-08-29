package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 3.594v9.969L9.719 9.28L8.28 10.72L13.562 16l-5.28 5.281l1.437 1.438L14 18.437v9.97l1.719-1.688l5-5l.687-.719l-.687-.719L16.437 16l4.282-4.281l.687-.719l-.687-.719l-5-5zm2 4.844L18.563 11L16 13.563zm0 10L18.563 21L16 23.563z"/>`),
		g.Group(children),
	)
}