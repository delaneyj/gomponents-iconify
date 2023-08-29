package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screenmaster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 9.2h3.7v0h25.9a3.983 3.983 0 0 1 3.7 3.7l.112 20.736M9.281 13.743L9.2 35.1a3.98 3.98 0 0 0 3.7 3.7h25.9v0h3.7"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.6 31.4v-3.7l11.1-11.1l3.7 3.7l-11.1 11.1Z"/>`),
		g.Group(children),
	)
}