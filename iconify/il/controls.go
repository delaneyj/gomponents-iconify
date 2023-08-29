package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Controls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M463 194H0v-93h463v93zm278-93v93h-92v81q0 5-3 8t-9 3h-93q-11 0-11-11V20q0-12 11-12h93q5 0 9 3t3 9v81h92zM0 564h255v93H0v-93zm440 0h301v93H440v81q0 5-3 8t-8 4h-93q-5 0-8-4t-4-8V483q0-5 4-8t8-3h93q11 0 11 11v81zM232 333h509v92H232v-92zM46 425H0v-92h46v-82q0-11 12-11h93q11 0 11 11v255q0 12-11 12H58q-12 0-12-12v-81z"/>`),
		g.Group(children),
	)
}