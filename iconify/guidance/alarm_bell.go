package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmBell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M21.753 11.48A12.56 12.56 0 0 1 15 20.228V23.5H4V20h9m-3.5-9.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm12.5 1a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm-12.5 6a8.5 8.5 0 1 1 0-17a8.5 8.5 0 0 1 0 17Z"/>`),
		g.Group(children),
	)
}