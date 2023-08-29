package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AwardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.25 9.5a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M5 9.5a7 7 0 1 1 12.923 3.732l2.727 4.723a.75.75 0 0 1-.796 1.11l-2.616-.52l-.858 2.526a.75.75 0 0 1-1.36.134l-2.72-4.711a7.119 7.119 0 0 1-.6 0l-2.72 4.711a.75.75 0 0 1-1.36-.132l-.861-2.52l-2.614.513a.75.75 0 0 1-.795-1.11l2.727-4.723A6.967 6.967 0 0 1 5 9.5Zm2.086 4.985a6.993 6.993 0 0 0 3.027 1.758l-1.607 2.783l-.54-1.579a.75.75 0 0 0-.854-.493l-1.637.322l1.61-2.79Zm6.801 1.758l1.605 2.779l.537-1.581a.75.75 0 0 1 .856-.495l1.638.326l-1.609-2.787a6.994 6.994 0 0 1-3.027 1.758ZM12 6.25a3.25 3.25 0 1 0 0 6.5a3.25 3.25 0 0 0 0-6.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}