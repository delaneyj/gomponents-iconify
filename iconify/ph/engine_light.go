package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EngineLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M248 106a6 6 0 0 0-6 6v26h-20v-18a14 14 0 0 0-14-14h-12.69a2 2 0 0 1-1.41-.58L158.58 70.1a13.9 13.9 0 0 0-9.89-4.1H126V46h26a6 6 0 0 0 0-12H88a6 6 0 0 0 0 12h26v20H48a14 14 0 0 0-14 14v58H14v-26a6 6 0 0 0-12 0v64a6 6 0 0 0 12 0v-26h20v22.69a13.9 13.9 0 0 0 4.1 9.89l35.32 35.32a13.9 13.9 0 0 0 9.89 4.1h65.38a13.9 13.9 0 0 0 9.89-4.1l35.32-35.32a2 2 0 0 1 1.41-.58H208a14 14 0 0 0 14-14v-18h20v26a6 6 0 0 0 12 0v-64a6 6 0 0 0-6-6Zm-38 62a2 2 0 0 1-2 2h-12.69a13.9 13.9 0 0 0-9.89 4.1l-35.32 35.32a2 2 0 0 1-1.41.58H83.31a2 2 0 0 1-1.41-.58L46.58 174.1a2 2 0 0 1-.58-1.41V80a2 2 0 0 1 2-2h100.69a2 2 0 0 1 1.41.58l35.32 35.32a13.9 13.9 0 0 0 9.89 4.1H208a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}