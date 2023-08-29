package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluelineconsole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 19.2h39v18h-39zm0 0v-8.4h16l6.9 8.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.21 22.74l5.64 5.46l-5.64 5.46"/>`),
		g.Group(children),
	)
}