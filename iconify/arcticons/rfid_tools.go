package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RfidTools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="29.736" cy="20.058" r="2.528" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.057 21.74a9.004 9.004 0 0 0-9.003-9.004"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.975 21.74A13.92 13.92 0 0 0 28.054 7.82m-10.601 6.872H8.01a3.51 3.51 0 0 0-3.51 3.51V36.67a3.51 3.51 0 0 0 3.51 3.51h31.98a3.51 3.51 0 0 0 3.51-3.51v-9.58"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.755 26.764a2.35 2.35 0 0 0-2.474-2.35a2.441 2.441 0 0 0-2.222 2.483v2.18a2.35 2.35 0 0 0 2.348 2.353h0a2.35 2.35 0 0 0 2.348-2.353h-2.348M8.11 31.43v-7.02h2.298a2.358 2.358 0 0 1 0 4.715H8.11m2.298.001l2.298 2.303m2.884.001v-7.02h2.298a2.358 2.358 0 0 1 0 4.715h-2.299m2.299.001l2.298 2.303"/>`),
		g.Group(children),
	)
}