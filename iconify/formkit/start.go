package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Start(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.5 10c-.28 0-.5-.22-.5-.5v-7c0-.28.22-.5.5-.5s.5.22.5.5v7c0 .28-.22.5-.5.5Z"/><path fill="currentColor" d="M7.5 15C3.92 15 1 12.18 1 8.72c0-2.41 1.46-4.64 3.72-5.68c.25-.11.55 0 .66.25c.11.25 0 .55-.25.66c-1.91.87-3.14 2.74-3.14 4.77c0 2.91 2.47 5.28 5.5 5.28s5.5-2.37 5.5-5.28c0-2.02-1.23-3.9-3.14-4.77a.502.502 0 0 1-.25-.66c.11-.25.41-.36.66-.25c2.26 1.03 3.72 3.26 3.72 5.68c0 3.46-2.92 6.28-6.5 6.28Z"/>`),
		g.Group(children),
	)
}