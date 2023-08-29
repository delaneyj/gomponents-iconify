package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 4.594v22.812l1.719-1.687l9-9l.687-.719l-.687-.719l-9-9zm2 4.843L20.563 16L14 22.563z"/>`),
		g.Group(children),
	)
}