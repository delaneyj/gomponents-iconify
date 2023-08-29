package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseMedicalThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 5h6a1 1 0 0 1 1 1v2h-8V6a1 1 0 0 1 1-1Zm-3 1v2H7.5A4.5 4.5 0 0 0 3 12.5v11A4.5 4.5 0 0 0 7.5 28h17a4.5 4.5 0 0 0 4.5-4.5v-11A4.5 4.5 0 0 0 24.5 8H22V6a3 3 0 0 0-3-3h-6a3 3 0 0 0-3 3Zm6 8a1 1 0 0 1 1 1v2h2a1 1 0 1 1 0 2h-2v2a1 1 0 1 1-2 0v-2h-2a1 1 0 1 1 0-2h2v-2a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}