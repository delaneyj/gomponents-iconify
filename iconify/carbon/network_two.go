package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 17h5.142a4 4 0 1 0 0-2H17V7h5.142a4 4 0 1 0 0-2H17a2.002 2.002 0 0 0-2 2v8H9.858a4 4 0 1 0 0 2H15v8a2.002 2.002 0 0 0 2 2h5.142a4 4 0 1 0 0-2H17Zm9-3a2 2 0 1 1-2 2a2.002 2.002 0 0 1 2-2Zm0-10a2 2 0 1 1-2 2a2.002 2.002 0 0 1 2-2ZM6 18a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Zm20 6a2 2 0 1 1-2 2a2.002 2.002 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}