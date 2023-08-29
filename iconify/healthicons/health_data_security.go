package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthDataSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 7a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3h3a3 3 0 0 1 3 3v15h-2V10a1 1 0 0 0-1-1h-3a3 3 0 0 1-3 3h-8a3 3 0 0 1-3-3h-3a1 1 0 0 0-1 1v28a1 1 0 0 0 1 1h12v2H14a3 3 0 0 1-3-3V10a3 3 0 0 1 3-3h3Zm3-1a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-8Zm3 10v3h-3v2h3v3h2v-3h3v-2h-3v-3h-2Zm5 11a1 1 0 0 0-1 1v6.789a7 7 0 0 0 3.117 5.824l3.328 2.22a1 1 0 0 0 1.11 0l3.328-2.22A7 7 0 0 0 41 34.79V28a1 1 0 0 0-1-1H28Zm9.707 5.707l-4 4a1 1 0 0 1-1.414 0l-2-2l1.414-1.414L33 34.586l3.293-3.293l1.414 1.414Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}