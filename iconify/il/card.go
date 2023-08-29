package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Card(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M718 1q10 0 17 6t6 17v510q0 10-6 16t-17 7H23q-10 0-16-7t-7-16V24Q0 14 7 7t16-6h695zM93 325h208v-69H93v69zm393 70H93v69h393v-69zM649 94h-70v69h70V94z"/>`),
		g.Group(children),
	)
}