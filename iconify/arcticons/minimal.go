package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minimal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="8.32" cy="36.81" r="3.82" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.75 23.34l2.57 2.57l5.07-5.07M4.71 9.87l2.58 2.57l5.07-5.07m4.22 2.57H43.5M16.58 23.38H43.5M16.58 36.81H43.5"/>`),
		g.Group(children),
	)
}