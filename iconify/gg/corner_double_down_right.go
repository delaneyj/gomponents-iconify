package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDoubleDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.6 7.68L7.638 2.741L2.701 7.704l1.418 1.41L6.522 6.7l-.014 6.036a4.8 4.8 0 0 0 4.788 4.812l5.928.014l-2.239 2.303l1.434 1.394l4.88-5.019l-5.019-4.88l-1.394 1.434l2.436 2.369l-6.02-.015a2.4 2.4 0 0 1-2.394-2.406l.014-5.9l2.268 2.256L12.6 7.68Z"/>`),
		g.Group(children),
	)
}