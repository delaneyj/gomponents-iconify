package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 5h13a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-4.5l.5 3h1v1H9v-1h1l.5-3H6a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3m5.5 13l-.5 3h3l-.5-3h-2M6 6a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h13a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H6M1 8a5 5 0 0 1 5-5h12v1H6a4 4 0 0 0-4 4v6H1V8Z"/>`),
		g.Group(children),
	)
}