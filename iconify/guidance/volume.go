package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M17 6.5a5.5 5.5 0 1 1 0 11m0-8a2.5 2.5 0 0 1 0 5m-2.5 7h-.25l-.523-.548A16 16 0 0 0 2.154 16H1.5V8h.654a16 16 0 0 0 11.573-4.952l.523-.548h.25v19Z"/>`),
		g.Group(children),
	)
}