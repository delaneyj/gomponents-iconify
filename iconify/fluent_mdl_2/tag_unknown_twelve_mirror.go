package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagUnknownTwelveMirror(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1365 427q0-36 25-61t61-25q18 0 33 7t27 18t18 27t7 34q0 35-25 60t-60 25q-18 0-33-6t-27-18t-19-27t-7-34zM853 1630l248-248q28 74 70 135l-318 318l-810-811L1067 0h810v683h-170V171h-606l-853 853l605 606zm683 418v-171h171v171h-171zm85-1195q71 0 133 27t108 73t74 109t27 133h-171q0-35-13-66t-37-54t-55-37t-66-14q-35 0-66 13t-54 37t-36 55t-14 66q0 39 19 70t47 60t62 59t61 66t48 84t19 109v64h-171v-64q0-39-19-70t-47-60t-62-59t-61-67t-48-83t-19-109q0-70 26-132t73-109t109-74t133-27z"/>`),
		g.Group(children),
	)
}