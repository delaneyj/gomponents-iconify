package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileMoveO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M672 768h736V544q0-9 5-17q3-7 14-12q10-3 18-2q10 2 17 7l384 353q6 6 7 10q3 6 3 13t-3 13q-1 5-7 11l-384 351q-6 6-17 8q-8 1-18-2q-9-4-14-11q-5-8-5-18v-224H672q-15-2-23-10q-9-8-9-22V800q0-14 9-23t23-9zM0 544v896q0 25 13 49q14 21 35 34q20 13 48 13h1088q27 0 49-13q21-13 34-34q13-22 13-49v-320h-128v288H128V640h416q27 0 49-13q21-13 34-34q13-22 13-49V128h512v544h128V96q0-28-13-48q-13-21-34-35q-24-13-49-13H544q-26 0-58 9q-27 8-59 25q-29 16-47 34L68 380q-22 23-34 47q-15 28-24 59q-10 30-10 58zm512-32H136q4-11 11-24q6-12 11-17l313-313q4-4 17-12q21-9 24-10v376z"/>`),
		g.Group(children),
	)
}