package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M35.28 10.464a2 2 0 0 1 .256 2.816L28.27 22H30a2 2 0 1 1 0 4h-4v2h4a2 2 0 1 1 0 4h-4v4a2 2 0 0 1-4 0v-4h-4a2 2 0 0 1 0-4h4v-2h-4a2 2 0 0 1 0-4h1.73l-7.266-8.72a2 2 0 1 1 3.072-2.56L24 20.876l8.464-10.156a2 2 0 0 1 2.816-.256Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}