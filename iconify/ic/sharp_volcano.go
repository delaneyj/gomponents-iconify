package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpVolcano(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 8h-7l-2 5H6l-4 9h20zm-5-7h2v4h-2zm3.121 4.468L18.95 2.64l1.414 1.414l-2.829 2.828zM7.64 4.05l1.414-1.414l2.828 2.828l-1.414 1.415z"/>`),
		g.Group(children),
	)
}