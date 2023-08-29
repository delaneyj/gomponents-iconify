package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anyconnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.932 2.96s-10.991 8.362-5.864 18.08C27.196 30.758 45 28.3 45 28.3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.75 21S19.2 6.616 30.4 15.5s.1 28.9.1 28.9"/>`),
		g.Group(children),
	)
}