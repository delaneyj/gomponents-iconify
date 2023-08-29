package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.8 22.6l-4.2-4.2L12 22h-1v-7.6L6.4 19L5 17.6l4.9-4.9l-8.5-8.5l1.4-1.4l18.4 18.4l-1.4 1.4ZM13 18.15L14.15 17L13 15.85v2.3Zm1.1-6.85l-1.4-1.4l2.2-2.2L13 5.85v4.35l-2-2V2h1l5.7 5.7l-3.6 3.6Z"/>`),
		g.Group(children),
	)
}