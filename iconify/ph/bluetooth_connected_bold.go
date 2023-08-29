package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothConnectedBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M191.2 166.4L140 128l51.2-38.4a12 12 0 0 0 0-19.2l-64-48A12 12 0 0 0 108 32v72L63.2 70.4a12 12 0 0 0-14.4 19.2L100 128l-51.2 38.4a12 12 0 1 0 14.4 19.2L108 152v72a12 12 0 0 0 19.2 9.6l64-48a12 12 0 0 0 0-19.2ZM132 56l32 24l-32 24Zm0 144v-48l32 24Zm-84-56a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm168-16a16 16 0 1 1-16-16a16 16 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}