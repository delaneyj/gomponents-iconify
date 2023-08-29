package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlablaAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.77 9.5L20.93 21a1.63 1.63 0 0 0 1.8 1.47l.49-.05m8.94-4.25a3.29 3.29 0 0 1-2.94 3.6h0a3.29 3.29 0 0 1-3.6-2.94l-.22-2.13a3.29 3.29 0 0 1 3-3.6h0A3.29 3.29 0 0 1 32 16m.49 5.44l-.87-8.67M10.55 18.2a3.27 3.27 0 0 1 2.94-3.6h0a3.27 3.27 0 0 1 3.6 2.94l.22 2.12a3.3 3.3 0 0 1-2.95 3.61h0a3.3 3.3 0 0 1-3.6-2.95m.33 3.28L9.77 10.51M25.51 24.4l1.15 11.46a1.65 1.65 0 0 0 1.8 1.47h.49m8.95-4.26a3.29 3.29 0 0 1-2.9 3.6h0a3.29 3.29 0 0 1-3.61-2.94l-.21-2.13a3.29 3.29 0 0 1 2.9-3.6h0a3.29 3.29 0 0 1 3.6 2.94m.55 5.4l-.88-8.67M16.28 33.1a3.29 3.29 0 0 1 2.95-3.6h0a3.29 3.29 0 0 1 3.6 2.94l.17 2.13a3.28 3.28 0 0 1-2.94 3.6h0a3.28 3.28 0 0 1-3.6-2.94m.37 3.27l-1.32-13.09"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}