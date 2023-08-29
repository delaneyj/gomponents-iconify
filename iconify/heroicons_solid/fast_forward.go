package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.555 5.168A1 1 0 0 0 3 6v8a1 1 0 0 0 1.555.832L10 11.202V14a1 1 0 0 0 1.555.832l6-4a1 1 0 0 0 0-1.664l-6-4A1 1 0 0 0 10 6v2.798l-5.445-3.63Z"/>`),
		g.Group(children),
	)
}