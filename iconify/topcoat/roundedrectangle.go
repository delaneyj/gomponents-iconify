package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Roundedrectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M11.5 4.5c-6.939 0-11 4.47-11 11v11c0 6.971 3.859 11 11 11h18c7.4 0 11-4.029 11-11v-11c0-6.97-4.061-11-11-11h-18z"/>`),
		g.Group(children),
	)
}