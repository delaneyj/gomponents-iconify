package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothDisabledRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 14.4l-3.9 3.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l4.2-4.2l-7.8-7.8q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-3.5-3.5l-2.9 2.9q-.15.15-.325.225T12 21.6q-.4 0-.7-.288t-.3-.737V14.4Zm2 3.75L14.15 17L13 15.85v2.3Zm1.1-6.85l-1.4-1.4l2.2-2.2L13 5.85v4.35l-2-2V3.425q0-.45.3-.737T12 2.4q.2 0 .375.075t.325.225L17 7q.15.15.213.325t.062.375q0 .2-.063.375T17 8.4l-2.9 2.9Z"/>`),
		g.Group(children),
	)
}