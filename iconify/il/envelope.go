package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M348 250q22 6 45 0l348-87v371q0 10-6 16t-17 7H23q-10 0-16-7t-7-16V163zM718 1q10 0 17 6t6 17v70l-348 87q-23 5-45 0L0 94V24Q0 14 7 7t16-6h695z"/>`),
		g.Group(children),
	)
}