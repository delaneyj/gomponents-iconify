package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M21 7a3 3 0 1 1 6 0v24a3 3 0 1 1-6 0V7Zm0 34a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z"/>`),
		g.Group(children),
	)
}