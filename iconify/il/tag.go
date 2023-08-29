package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M725 348q7 7 7 16t-7 16L398 708q-29 27-66 27H23q-10 0-16-7t-7-16V402q0-38 27-65L354 10q7-7 17-7t16 7zM185 619q29 0 50-21t20-49t-20-49t-50-20t-49 20t-20 49t20 49t49 21z"/>`),
		g.Group(children),
	)
}