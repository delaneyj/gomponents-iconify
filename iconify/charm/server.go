package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Server(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 3.25h12.5v10H1.75zm.5 5h11.5m-9 2.5v0m0-5v0m6.5 0h-3m3 5h-3"/>`),
		g.Group(children),
	)
}