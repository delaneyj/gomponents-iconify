package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gpsstatus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.8 25.51a1.78 1.78 0 1 1 0 2.52a1.78 1.78 0 0 1 0-2.52ZM10.56 37a9.52 9.52 0 0 1 .12-13.44a1.47 1.47 0 0 1 2.08.09l.12.14l11.35 10.93a1.71 1.71 0 0 1 0 2.42a9.58 9.58 0 0 1-13.55 0a.93.93 0 0 1-.14-.14Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.85 43.6a25.77 25.77 0 0 0 .2-3.19A25.46 25.46 0 0 0 7.59 15a25.77 25.77 0 0 0-3.19.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.89 37.3c.07-.81.11-1.62.11-2.45A27.85 27.85 0 0 0 13.15 7c-.83 0-1.64 0-2.45.11"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}