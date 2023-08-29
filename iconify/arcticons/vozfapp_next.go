package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VozfappNext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.634 4.5l13.744 39l8.63-8.608l-7.458-22.636L9.633 4.5Zm22.375 30.392l6.357-6.483l-1.99-5.526l-10.077-5.32"/>`),
		g.Group(children),
	)
}