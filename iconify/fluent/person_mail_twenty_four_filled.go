package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonMailTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 15.5c0-.563.186-1.082.5-1.5H5.253a2.249 2.249 0 0 0-2.25 2.25v.919c0 .572.18 1.13.511 1.596C5.056 20.929 7.58 22 11 22h.05a2.514 2.514 0 0 1-.05-.5v-6Zm0-13.495a5 5 0 1 1 0 10a5 5 0 0 1 0-10Zm6.51 16.922l-5.49-3.203A2 2 0 0 1 14 14h7a2 2 0 0 1 2 1.97l-5.49 2.957Zm.227 1.014l5.264-2.834V21a2 2 0 0 1-2 2h-7a2 2 0 0 1-2-2v-4.13l5.248 3.062a.5.5 0 0 0 .489.009Z"/>`),
		g.Group(children),
	)
}