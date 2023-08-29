package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 4H10a2.002 2.002 0 0 0-2 2v14.556A3.955 3.955 0 0 0 6 20a4 4 0 1 0 4 4V12h15v8.556A3.954 3.954 0 0 0 23 20a4 4 0 1 0 4 4V6a2.002 2.002 0 0 0-2-2ZM6 26a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Zm17 0a2 2 0 1 1 2-2a2.003 2.003 0 0 1-2 2ZM10 6h15v4H10Z"/>`),
		g.Group(children),
	)
}