package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watertracker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5c0 3.945-12 13.717-12 27.405A11.803 11.803 0 0 0 24 43.5a11.803 11.803 0 0 0 12-11.595C36 18.217 24 8.445 24 4.5Z"/>`),
		g.Group(children),
	)
}