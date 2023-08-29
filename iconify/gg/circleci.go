package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circleci(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12S18.627 0 12 0ZM4.144 13.517a8 8 0 1 0-.006-3l2.59-.01a5.478 5.478 0 1 1 .004 3l-2.588.01ZM9.522 12a2.478 2.478 0 1 0 4.956 0a2.478 2.478 0 0 0-4.956 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}