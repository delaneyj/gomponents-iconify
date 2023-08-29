package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yamusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.91 4.5v17.85a7.57 7.57 0 1 0 2.24 5.38V13.61l11 3.4V9ZM23 12h-.07a15.77 15.77 0 1 0 15.44 15.72"/>`),
		g.Group(children),
	)
}