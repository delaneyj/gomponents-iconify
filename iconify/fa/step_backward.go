package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M979 13q19-19 32-13t13 32v1472q0 26-13 32t-32-13L269 813q-9-9-13-19v678q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V64q0-26 19-45T64 0h128q26 0 45 19t19 45v678q4-10 13-19z"/>`),
		g.Group(children),
	)
}