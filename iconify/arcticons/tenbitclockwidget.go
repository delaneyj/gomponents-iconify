package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tenbitclockwidget(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="26.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="16.4" ry="16.6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.1 28.9V17.5M19 4.5h10.1v3.4H19zm15.6 9.6l2.9-2.7l2.1 2.3l-2.9 2.7"/>`),
		g.Group(children),
	)
}