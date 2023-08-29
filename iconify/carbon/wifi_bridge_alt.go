package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiBridgeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 26h-2a5.006 5.006 0 0 0-5-5v-2a7.008 7.008 0 0 1 7 7Z"/><path fill="currentColor" d="M18 26h-2A10.011 10.011 0 0 0 6 16v-2a12.014 12.014 0 0 1 12 12zm8-13a7.008 7.008 0 0 1-7-7h2a5.006 5.006 0 0 0 5 5z"/><path fill="currentColor" d="M26 18A12.014 12.014 0 0 1 14 6h2a10.011 10.011 0 0 0 10 10zM7.707 24.293a1 1 0 0 0-1.414 0L2 28.586L3.414 30l4.293-4.293a1 1 0 0 0 0-1.414zM28.586 2l-4.293 4.293a1 1 0 0 0 1.414 1.414L30 3.414z"/>`),
		g.Group(children),
	)
}