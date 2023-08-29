package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 21h2V11h6v2h8V5h-7V3H4v18Zm8-16H6v4h7v2h5V7h-6V5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}