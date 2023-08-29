package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoryOfMe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="18.332" cy="15.02" r="5.389" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="29.668" cy="15.02" r="5.389" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.056" cy="9.889" r="5.389" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.056 11.267V43.5m0-9.298h-4.422m4.422 2.935h-4.422m4.422 2.89h-4.422"/>`),
		g.Group(children),
	)
}