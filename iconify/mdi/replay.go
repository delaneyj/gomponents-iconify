package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Replay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 5V1L7 6l5 5V7a6 6 0 0 1 6 6a6 6 0 0 1-6 6a6 6 0 0 1-6-6H4a8 8 0 0 0 8 8a8 8 0 0 0 8-8a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}