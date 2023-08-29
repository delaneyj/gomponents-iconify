package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothConnectedThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M186.4 172.8L126.67 128l59.73-44.8a4 4 0 0 0 0-6.4l-64-48A4 4 0 0 0 116 32v88L58.4 76.8a4 4 0 0 0-4.8 6.4l59.73 44.8l-59.73 44.8a4 4 0 0 0 4.8 6.4L116 136v88a4 4 0 0 0 6.4 3.2l64-48a4 4 0 0 0 0-6.4ZM124 40l53.33 40L124 120Zm0 176v-80l53.33 40Zm-72-80a8 8 0 1 1 8-8a8 8 0 0 1-8 8Zm152-8a8 8 0 1 1-8-8a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}