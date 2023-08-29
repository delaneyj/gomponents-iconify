package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CatchUpTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.278 8.157a.25.25 0 0 1 .462-.001l3.236 7.766c.574 1.377 2.497 1.451 3.175.123l1.6-3.136a.75.75 0 0 1 .668-.41h.644a2 2 0 1 0 .205-1.5h-.85a2.25 2.25 0 0 0-2.004 1.228l-1.6 3.136a.25.25 0 0 1-.453-.018L11.124 7.58c-.6-1.441-2.644-1.434-3.235.011l-1.202 2.943a.75.75 0 0 1-.695.467h-1.26a2 2 0 1 0 .205 1.5h1.055a2.25 2.25 0 0 0 2.083-1.4l1.203-2.943Z"/>`),
		g.Group(children),
	)
}