package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ericsson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.717 5.723a2 2 0 1 0 1.69 3.625l10.876-5.071a2 2 0 0 0-1.69-3.625L7.717 5.723ZM4.75 15.381a2 2 0 0 1 .967-2.658l10.876-5.071a2 2 0 1 1 1.69 3.625L7.407 16.348a2 2 0 0 1-2.657-.967Zm-2 7a2 2 0 0 1 .967-2.658l10.876-5.071a2 2 0 1 1 1.69 3.625L5.407 23.348a2 2 0 0 1-2.657-.967Z"/>`),
		g.Group(children),
	)
}