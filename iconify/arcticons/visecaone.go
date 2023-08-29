package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Visecaone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.34 36.47a3.36 3.36 0 0 1-2.38 1h0a3.37 3.37 0 0 1-3.37-3.38v-2.23A3.37 3.37 0 0 1 34 28.48h0a3.37 3.37 0 0 1 3.37 3.38v1m-.04.11h-6.75"/><rect width="6.75" height="8.98" x="10.67" y="28.48" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.62 31.86A3.38 3.38 0 0 1 24 28.48h0a3.37 3.37 0 0 1 3.37 3.38v5.6m-6.75-8.98v8.98"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><circle cx="32.54" cy="15.46" r="4.79" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}