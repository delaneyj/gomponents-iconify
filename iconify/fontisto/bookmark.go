package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.08 1.6h8.88a.516.516 0 0 1 .48.478V20.24L7.6 16.56l-1.12-1.04l-1.12 1.04l-3.76 3.68V2.08a.516.516 0 0 1 .478-.48h.002zm0-1.6A2.103 2.103 0 0 0 0 1.995V24l6.56-6.24L13.12 24V2.08A2.119 2.119 0 0 0 11.042 0h-.002z"/>`),
		g.Group(children),
	)
}