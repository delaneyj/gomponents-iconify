package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M21.067 10.8l-6.667-5C13.411 5.058 12 5.764 12 7v3L6.4 5.8C5.411 5.058 4 5.764 4 7v10c0 1.236 1.411 1.942 2.4 1.2L12 14v3c0 1.236 1.411 1.942 2.4 1.2l6.667-5c.8-.6.8-1.8 0-2.4z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}