package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Police(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 24h2v6h-2z"/><path fill="currentColor" d="M21 20H7a3.003 3.003 0 0 1-3-3V6H2v11a5.006 5.006 0 0 0 5 5h14a4.942 4.942 0 0 1 2.105.481L17 28.586L18.414 30l6.307-6.307A4.96 4.96 0 0 1 26 27v3h2v-3a7.008 7.008 0 0 0-7-7Z"/><path fill="currentColor" d="m25.275 4.039l-7-2a1.001 1.001 0 0 0-.55 0l-7 2a1 1 0 0 0-.695 1.203L11 9.123V11a7 7 0 1 0 14 0V9.123l.97-3.88a1 1 0 0 0-.695-1.204ZM18 4.04l5.795 1.656L23.22 8H19V6h-2v2h-4.219l-.576-2.304ZM18 16a5.006 5.006 0 0 1-5-5v-1h10v1a5.006 5.006 0 0 1-5 5Z"/>`),
		g.Group(children),
	)
}