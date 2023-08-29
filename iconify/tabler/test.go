package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Test(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M3 9h18M3 9v6m0-6V5a2 2 0 0 1 2-2h4m12 6v6m0-6V5a2 2 0 0 0-2-2h-4M3 15v4a2 2 0 0 0 2 2h4m-6-6h18m0 0v4a2 2 0 0 1-2 2h-4M9 3v18M9 3h6M9 21h6m0-18v18" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}