package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rdclient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.65 20.05V6.85H4.55v22h14.9V33c0 1 .1 2-2.4 2.1h-2.9a2.15 2.15 0 0 0-2.1 2.1h12"/><circle cx="32.45" cy="30.15" r="11" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m37 24l-3.9 3.9l3.9 3.9m-8.95-3.25l3.9 3.9l-3.9 3.9m-8.6-7.5h2"/>`),
		g.Group(children),
	)
}