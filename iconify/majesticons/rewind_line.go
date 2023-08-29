package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10 8v8l-5.333-4L10 8zm2-1c0-1.236-1.411-1.942-2.4-1.2l-6.667 5c-.8.6-.8 1.8 0 2.4l6.667 5c.989.742 2.4.036 2.4-1.2v-3l5.6 4.2c.989.742 2.4.036 2.4-1.2V7c0-1.236-1.411-1.942-2.4-1.2L12 10V7zm.667 5L18 8v8l-5.333-4z"/></g>`),
		g.Group(children),
	)
}