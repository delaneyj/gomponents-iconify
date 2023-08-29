package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 20L1 12l10-8v5c5.523 0 10 4.477 10 10c0 .273-.01.543-.032.81A8.999 8.999 0 0 0 13 15h-2v5Z"/>`),
		g.Group(children),
	)
}