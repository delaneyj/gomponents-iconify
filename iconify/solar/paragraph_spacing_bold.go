package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphSpacingBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.25 3A.75.75 0 0 1 4 2.25h16a.75.75 0 0 1 0 1.5H4A.75.75 0 0 1 3.25 3Zm0 18a.75.75 0 0 1 .75-.75h16a.75.75 0 0 1 0 1.5H4a.75.75 0 0 1-.75-.75Zm9.28-16.03a.75.75 0 0 0-1.06 0l-3 3a.75.75 0 0 0 1.06 1.06l1.72-1.72v9.38l-1.72-1.72a.75.75 0 0 0-1.06 1.06l3 3a.75.75 0 0 0 1.06 0l3-3a.75.75 0 1 0-1.06-1.06l-1.72 1.72V7.31l1.72 1.72a.75.75 0 1 0 1.06-1.06l-3-3Z"/>`),
		g.Group(children),
	)
}