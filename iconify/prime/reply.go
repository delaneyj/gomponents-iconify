package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.55 15.59a.75.75 0 0 1-.55-1.28l3.92-3.89L14 6.53a.75.75 0 0 1 0-1.06a.75.75 0 0 1 1.06 0l4.46 4.42a.75.75 0 0 1 0 1.06l-4.46 4.42a.7.7 0 0 1-.51.22Z"/><path fill="currentColor" d="M5 18.75a.76.76 0 0 1-.75-.75v-7.58A.76.76 0 0 1 5 9.67h14a.75.75 0 0 1 0 1.5H5.75V18a.76.76 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}