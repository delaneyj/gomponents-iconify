package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.5 7.5h5a3 3 0 1 0 0-6H3.58a.08.08 0 0 0-.08.08V7.5Zm0 0h6a3 3 0 1 1 0 6H3.59a.09.09 0 0 1-.09-.09V7.5Z"/>`),
		g.Group(children),
	)
}