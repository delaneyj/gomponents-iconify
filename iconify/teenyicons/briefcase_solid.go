package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 2.5V3H1.5A1.5 1.5 0 0 0 0 4.5V11h15V4.5A1.5 1.5 0 0 0 13.5 3H10v-.5a2.5 2.5 0 0 0-5 0ZM7.5 1A1.5 1.5 0 0 0 6 2.5V3h3v-.5A1.5 1.5 0 0 0 7.5 1Z" clip-rule="evenodd"/><path fill="currentColor" d="M15 12H0v1.5A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5V12Z"/>`),
		g.Group(children),
	)
}