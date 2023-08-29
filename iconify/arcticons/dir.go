package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dir(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="40" height="32" x="4" y="8" fill="none" stroke="currentColor" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M15.79 14.9h2.81v18.2h-2.81zm5.62 0c6 0 10.8 4.07 10.8 9.1c0 5-4.84 9.1-10.8 9.1V14.9Z"/>`),
		g.Group(children),
	)
}