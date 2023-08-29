package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForecastLightningThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 28a10 10 0 0 1 0-20h4v5l6-6l-6-6v5h-4a12 12 0 0 0 0 24Z"/><path fill="currentColor" d="m11.67 24l-1.736-1l2.287-4H8.332l3.993-7l1.737 1l-2.284 4h3.89l-3.998 7zM20 20h-4v2h4v2h-3v2h3v2h-4v2h4a2.003 2.003 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2zm8 10h-2a2.002 2.002 0 0 1-2-2v-6a2.002 2.002 0 0 1 2-2h2a2.002 2.002 0 0 1 2 2v6a2.002 2.002 0 0 1-2 2zm-2-8v6h2v-6z"/>`),
		g.Group(children),
	)
}