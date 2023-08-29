package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineageUpdater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="23.478" height="39" x="12.261" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.062"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.261 8.88h23.478M12.261 39.12h23.478M24 30V18m3.65 8.35L24 30l-3.65-3.65"/>`),
		g.Group(children),
	)
}