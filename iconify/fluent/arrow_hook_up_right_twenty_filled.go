package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowHookUpRightTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 6c.06 0-.06-.002 0 0c.023-.002.227 0 .25 0h4.393l-2.268-2.268a.75.75 0 1 1 1.061-1.06l3.353 3.352a.749.749 0 0 1 .212.639a.747.747 0 0 1-.215.443l-3.247 3.248a.75.75 0 1 1-1.06-1.061L13.27 7.5H9a3.5 3.5 0 1 0 0 7h4.25a.75.75 0 0 1 0 1.5H9A5 5 0 0 1 9 6Z"/>`),
		g.Group(children),
	)
}