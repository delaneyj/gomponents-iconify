package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func W(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m590 487l156-362h79L591 667L412 261L234 667L0 125h79l156 362L412 82z"/>`),
		g.Group(children),
	)
}