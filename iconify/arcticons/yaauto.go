package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yaauto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.14 11a5.91 5.91 0 0 1 9.72 0L42.5 30.71a5.62 5.62 0 0 1-7.33 8.13l-7.31-4a8 8 0 0 0-7.71 0l-7.32 4a5.62 5.62 0 0 1-7.33-8.13Z"/>`),
		g.Group(children),
	)
}