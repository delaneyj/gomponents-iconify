package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jelly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5a21.5 21.5 0 0 0 0 43c4.45 0 8.06-9.63 8.06-21.5S28.45 2.5 24 2.5Zm0 43a21.49 21.49 0 0 0 1.09-42.95c8 .81 14.25 10.17 14.25 21.45c0 11.87-6.87 21.5-15.34 21.5Z"/>`),
		g.Group(children),
	)
}