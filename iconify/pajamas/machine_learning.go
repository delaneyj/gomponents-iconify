package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MachineLearning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.75 3.5a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5ZM9.5 8a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM8 11a3.001 3.001 0 0 0 2.905-2.25h1.845c.071.095.155.179.25.25v3.75a1.25 1.25 0 1 0 1.5 0V9a1.25 1.25 0 1 0-1.75-1.75h-1.845A3.005 3.005 0 0 0 8.75 5.095V3.25A1.25 1.25 0 1 0 7 1.5H3.25a1.25 1.25 0 1 0 0 1.5H7c.071.095.155.179.25.25v1.845A3.001 3.001 0 0 0 8 11Zm-5.75 4a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Zm7-1.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0ZM3.5 8A1.25 1.25 0 1 1 1 8a1.25 1.25 0 0 1 2.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}