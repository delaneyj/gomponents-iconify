package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothSlashBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m216.88 207.93l-160-176a12 12 0 1 0-17.76 16.14L107 122.75L48.8 166.4a12 12 0 1 0 14.4 19.2L108 152v72a12 12 0 0 0 19.2 9.6l47.91-35.94l24 26.41a12 12 0 0 0 17.76-16.14ZM132 200v-48l5 3.77l21.87 24.06ZM108 59.74V32a12 12 0 0 1 19.2-9.6l64 48a12 12 0 0 1 0 19.2l-27.1 20.33a12 12 0 0 1-14.4-19.2L164 80l-32-24v3.74a12 12 0 0 1-24 0Z"/>`),
		g.Group(children),
	)
}