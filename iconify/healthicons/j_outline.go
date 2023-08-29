package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M27 12a3 3 0 1 1 6 0v18a9 9 0 1 1-18 0a3 3 0 0 1 6 0a3 3 0 1 0 6 0V12Zm3-1a1 1 0 0 0-1 1v18a5 5 0 0 1-10 0a1 1 0 1 0-2 0a7 7 0 1 0 14 0V12a1 1 0 0 0-1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}