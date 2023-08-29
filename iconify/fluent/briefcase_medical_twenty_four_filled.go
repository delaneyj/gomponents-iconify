package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseMedicalTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.25 3.5h3.5a.75.75 0 0 1 .75.75V6h-5V4.25a.75.75 0 0 1 .75-.75ZM8 4.25V6H6.25A3.25 3.25 0 0 0 3 9.25v7.5A3.25 3.25 0 0 0 6.25 20h11.5A3.25 3.25 0 0 0 21 16.75v-7.5A3.25 3.25 0 0 0 17.75 6H16V4.25A2.25 2.25 0 0 0 13.75 2h-3.5A2.25 2.25 0 0 0 8 4.25Zm3.5 6.5a.75.75 0 0 1 1.5 0v1.75h1.75a.75.75 0 0 1 0 1.5H13v1.75a.75.75 0 0 1-1.5 0V14H9.75a.75.75 0 0 1 0-1.5h1.75v-1.75Z"/>`),
		g.Group(children),
	)
}