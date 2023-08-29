package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freetrade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.081 4.5l6.03 3.433l.071 6.075a4.434 4.434 0 0 1-2.352 3.53l-9.503 5.601a4.054 4.054 0 0 0-2.088 3.598l-.07 4.595l-6.228-3.448l-.033-4.18a4.27 4.27 0 0 1 2.136-4.075L19.65 13.15a3.763 3.763 0 0 0 1.396-3.22ZM31.39 19.476a1.81 1.81 0 0 0-1.55.175L19.783 25.53a4.27 4.27 0 0 0-2.137 4.076l.033 4.178l6.23 3.447l.07-4.595a4.052 4.052 0 0 1 2.087-3.596l9.883-5.659c.775-.488.88-.913.33-1.216Zm3.97 10.689a1.32 1.32 0 0 0-1.344.033l-2.813 1.6a4.27 4.27 0 0 0-2.137 4.077l.033 4.178l6.23 3.447l.07-4.595a4.052 4.052 0 0 1 2.087-3.596l3.607-1.99l-5.733-3.154Z"/>`),
		g.Group(children),
	)
}