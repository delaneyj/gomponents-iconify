package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalligraphyPenTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 2.75a.75.75 0 0 0-1.5 0v3a1.75 1.75 0 0 0 1.543 1.738L6.527 9.993a3.868 3.868 0 0 0 .119 3.143l3.99 7.95c.149.298.363.535.614.693V12.3a1.5 1.5 0 1 1 1.5 0v9.48c.251-.158.465-.395.615-.692l3.99-7.951c.481-.96.526-2.137.118-3.143l-1.016-2.505A1.75 1.75 0 0 0 18 5.75v-3a.75.75 0 0 0-1.5 0v3a.25.25 0 0 1-.25.25h-8.5a.25.25 0 0 1-.25-.25v-3Z"/>`),
		g.Group(children),
	)
}