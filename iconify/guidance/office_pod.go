package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OfficePod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M8.5 15V9.5h-.676a3 3 0 0 0-2.905 2.25L3.5 17.25v.25H11v1.456c0 1.24 0 2.044.554 3.305c0 0 .62 1.239 1.446 1.239m-4 0H3.5a3 3 0 0 1-3-3v-17a3 3 0 0 1 3-3h17a3 3 0 0 1 3 3v17a3 3 0 0 1-3 3h-4v-9m-4.5 0h9m-12.613-7s-1.201-.75-1.201-1.688a1.313 1.313 0 0 1 2.625 0C9.81 6.75 8.613 7.5 8.613 7.5h-.226Z"/>`),
		g.Group(children),
	)
}