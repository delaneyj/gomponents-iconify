package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransferAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.924 10.383a1 1 0 0 0-.217-1.09l-5-5l-1.414 1.414L16.586 9H4v2h15a1 1 0 0 0 .924-.617zM4.076 13.617a1 1 0 0 0 .217 1.09l5 5l1.414-1.414L7.414 15H20v-2H5a.999.999 0 0 0-.924.617z"/>`),
		g.Group(children),
	)
}