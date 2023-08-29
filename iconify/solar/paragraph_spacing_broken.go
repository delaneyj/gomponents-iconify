package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphSpacingBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 21h9m7 0h-3M4 3h16m-8 2.5l3 3m-3-3l-3 3m3-3v13m0 0l3-3m-3 3l-3-3"/>`),
		g.Group(children),
	)
}