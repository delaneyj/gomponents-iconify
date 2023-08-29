package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Escalator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 22.5h5.03a8 8 0 0 0 6.813-3.807L18 9.5h5.5v-8h-5.03a8 8 0 0 0-6.813 3.807L6 14.5H.5v8Z"/>`),
		g.Group(children),
	)
}