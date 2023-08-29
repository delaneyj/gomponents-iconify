package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m3.594 12l1.687 1.72l10 10l.72.686l.72-.687l10-10L28.405 12H3.594zm4.844 2h15.124L16 21.563L8.437 14z"/>`),
		g.Group(children),
	)
}