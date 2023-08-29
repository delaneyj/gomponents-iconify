package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Peertube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.83 4.5l16.89 9.75L42.6 24l-16.88 9.75L8.83 43.5V24l8.36 4.83l8.42 4.87V14.27l-8.42 4.87L8.83 24V4.5Z"/>`),
		g.Group(children),
	)
}