package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 4.504H7v14.992h6a4 4 0 0 0 .604-7.955A4 4 0 0 0 11 4.505Zm-2 2h2a2 2 0 1 1 0 4H9v-4Zm0 10.991v-4h4a2 2 0 1 1 0 4H9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}