package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.08.858h15.84a1 1 0 0 1 .78 1.625l-6.48 8.101a1 1 0 0 0-.22.625v5.149a1 1 0 0 1-.4.8l-2 1.5a1 1 0 0 1-1.6-.8v-6.65a1 1 0 0 0-.22-.624L.3 2.483A1 1 0 0 1 1.08.858z"/>`),
		g.Group(children),
	)
}