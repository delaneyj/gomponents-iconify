package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rotate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 14v6a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2zM13.914 2.914L11.828 5H14a8 8 0 0 1 8 8h-2c0-3.308-2.692-6-6-6h-2.172l2.086 2.086L12.5 10.5L8 6l1.414-1.414L12.5 1.5l1.414 1.414z"/>`),
		g.Group(children),
	)
}