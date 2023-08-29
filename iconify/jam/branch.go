package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Branch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm1-7.002v3.173a3.001 3.001 0 1 1-2 0V5.829a3.001 3.001 0 1 1 2 0v2.34c.312-.11.647-.17.997-.171l6.037-.006a1 1 0 0 0 .999-1V5.84A3.001 3.001 0 0 1 13 0a3 3 0 0 1 1.033 5.817v1.175a3 3 0 0 1-2.997 3l-6.037.006a1 1 0 0 0-.999 1zM3 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm10 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2z"/>`),
		g.Group(children),
	)
}