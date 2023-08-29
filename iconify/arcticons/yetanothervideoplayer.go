package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yetanothervideoplayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.75 7.11h-19.5L4.5 24l9.75 16.89h19.5L43.5 24ZM19.12 30.82V17.17l13.65 6.88Z"/>`),
		g.Group(children),
	)
}