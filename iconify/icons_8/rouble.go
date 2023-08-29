package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 6v14H7v2h2v4h2v-4h6v-2h-6v-3h7.5c3.026 0 5.5-2.474 5.5-5.5S21.526 6 18.5 6H9zm2 2h7.5c1.944 0 3.5 1.556 3.5 3.5S20.444 15 18.5 15H11V8z"/>`),
		g.Group(children),
	)
}