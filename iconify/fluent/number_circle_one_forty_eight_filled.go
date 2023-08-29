package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleOneFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Zm2.5-29.75v18.5a1.25 1.25 0 1 1-2.5 0V19.672c-1.006 1.297-2.406 2.546-4.318 3.25a1.25 1.25 0 1 1-.864-2.345c1.98-.73 3.286-2.316 4.117-3.833c.41-.75.687-1.454.861-1.968a10.334 10.334 0 0 0 .224-.755l.008-.034v-.004l.001-.001a1.25 1.25 0 0 1 2.471.268Z"/>`),
		g.Group(children),
	)
}