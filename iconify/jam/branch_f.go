package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 11.856v3.173a3.001 3.001 0 1 1-2 0V6.687a3.001 3.001 0 1 1 2 0v2.34c.312-.11.647-.17.997-.171l6.037-.006a1 1 0 0 0 .999-1V6.699A3.001 3.001 0 0 1 13 .858a3 3 0 0 1 1.033 5.817V7.85a3 3 0 0 1-2.997 3l-6.037.006a1 1 0 0 0-.999 1z"/>`),
		g.Group(children),
	)
}