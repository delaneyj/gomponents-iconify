package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skysafari(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.85 3.41A21.49 21.49 0 0 1 8.05 38.5a21.5 21.5 0 1 0 21.8-35.09Z"/>`),
		g.Group(children),
	)
}