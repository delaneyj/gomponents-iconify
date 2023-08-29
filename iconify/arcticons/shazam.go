package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shazam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.55 6a3.84 3.84 0 0 1 2 7.18l-4.35 2.51a6.12 6.12 0 1 0 6.13 10.6l4.34-2.51a3.84 3.84 0 0 1 3.84 6.65l-4.35 2.51A13.79 13.79 0 1 1 12.38 9l4.35-2.51A3.75 3.75 0 0 1 18.55 6Zm9.79 7.26A13.78 13.78 0 0 1 35.62 39l-4.35 2.51a3.84 3.84 0 1 1-4-6.56l.15-.08l4.35-2.51a6.13 6.13 0 1 0-6.13-10.61l-4.35 2.51a3.83 3.83 0 0 1-3.83-6.64l4.35-2.51a13.73 13.73 0 0 1 6.51-1.83Z"/>`),
		g.Group(children),
	)
}