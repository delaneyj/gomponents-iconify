package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangularRuler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ffce31" d="M.7 61.8C-.5 63-.1 64 1.6 64h59.3c1.7 0 3.1-1.4 3.1-3.1V1.6c0-1.7-1-2.1-2.2-.9L.7 61.8m49.1-15c0 1.7-1.4 3.1-3.1 3.1h-11c-1.7 0-2.1-1-.9-2.2l12.8-12.8c1.2-1.2 2.2-.8 2.2.9v11"/><path fill="#89664c" d="M2.4 61.3h1V64h-1zm2.9 0h1V64h-1zm3 0h1V64h-1zm3-2.3h1v5h-1zm3.1 2.3h1V64h-1zm3 0h1V64h-1zm3 0h1V64h-1zm3-2.3h1v5h-1zm3.1 2.3h1V64h-1zm3 0h1V64h-1zm3 0h1V64h-1zm2.9-2.3h1v5h-1zm3.2 2.3h1V64h-1zm2.9 0h1V64h-1zm3 0h1V64h-1zm3-2.3h1v5h-1zm3.1 2.3h1V64h-1zm3 0h1V64h-1zm3 0h1V64h-1zm3-2.3h1v5h-1z"/>`),
		g.Group(children),
	)
}