package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nuclear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.45 24a6.45 6.45 0 0 0-3.225-5.586L34.75 5.38A21.5 21.5 0 0 1 45.5 24H30.45Zm-9.675 5.586a6.45 6.45 0 0 0 6.45 0L34.75 42.62a21.5 21.5 0 0 1-21.5 0l7.525-13.034Zm0-11.172A6.45 6.45 0 0 0 17.55 24H2.5A21.5 21.5 0 0 1 13.25 5.38l7.525 13.034Z"/>`),
		g.Group(children),
	)
}