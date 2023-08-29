package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyLinkManager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="30.766" height="30.766" x="11.991" y="4.626" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3" ry="3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.757 12.082v26.544a3 3 0 0 0 3 3h26.544m3.307-32.851h-9.081v15.749l4.541-2.811l4.54 2.811V8.775z"/>`),
		g.Group(children),
	)
}