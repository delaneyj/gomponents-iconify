package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpCommentBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2v20l4-4h16V2H2zm17 11l-2.5-1.5L14 13V5h5v8z"/>`),
		g.Group(children),
	)
}