package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPieOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.25 3.784a8.25 8.25 0 1 0 6.669 13.964l-6.39-5.165A.75.75 0 0 1 11.25 12V3.784Zm1.5 0v7.466h7.466a8.252 8.252 0 0 0-7.466-7.466Zm7.466 8.966h-6.095l4.741 3.831a8.198 8.198 0 0 0 1.354-3.831ZM2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75a9.712 9.712 0 0 1-2.167 6.13A9.733 9.733 0 0 1 12 21.75A9.75 9.75 0 0 1 2.25 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}