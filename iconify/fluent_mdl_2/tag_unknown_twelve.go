package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagUnknownTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1365 427q0-36 25-61t61-25q18 0 33 7t27 18t18 27t7 34q0 35-25 60t-60 25q-18 0-33-6t-27-18t-19-27t-7-34zM853 1630l248-248q28 74 70 135l-318 318l-810-811L1067 0h810v683h-170V171h-606l-853 853l605 606zm683 418v-171h171v171h-171zm85-1195q71 0 133 27t108 73t74 109t27 133q0 61-19 108t-50 87t-69 76t-79 76q-17 17-25 36t-12 41t-3 44t1 44h-171v-64q0-62 19-109t47-84t62-67t61-59t48-59t19-70q0-35-13-66t-37-54t-55-37t-66-14q-35 0-66 13t-54 37t-36 55t-14 66h-171q0-70 26-132t73-109t109-74t133-27z"/>`),
		g.Group(children),
	)
}