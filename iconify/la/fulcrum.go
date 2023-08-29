package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fulcrum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m15 0l-1.28 12.2L10.1 16l3.62 3.8L15 32V19l-3-3l3-3V0zm2 0v13l3 3l-3 3v13l1.28-12.2L21.9 16l-3.62-3.8L17 0zM9.5 7L7 10l2.5 3l2.5-3l-2.5-3zm13 0L20 10l2.5 3l2.5-3l-2.5-3z"/>`),
		g.Group(children),
	)
}