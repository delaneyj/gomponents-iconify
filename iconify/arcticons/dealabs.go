package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dealabs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.23 35.05l3.926 7.45h27.939l3.675-7.45l-13.491-19.885V5.5h-8.004v9.514L6.23 35.05z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.376 30.048l-6.624 3.189l-19.661-6.549"/>`),
		g.Group(children),
	)
}