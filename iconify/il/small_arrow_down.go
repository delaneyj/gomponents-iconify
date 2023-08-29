package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmallArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M371 148q0 9-7 16L202 326q-7 7-16 7t-17-7L7 164q-7-7-7-17t7-16t16-7t16 7l146 146l146-146q7-7 17-7t16 7q7 9 7 17z"/>`),
		g.Group(children),
	)
}