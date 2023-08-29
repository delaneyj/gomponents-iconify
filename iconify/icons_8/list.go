package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func List(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 5v6h6V5H4zm2 2h2v2H6V7zm6 0v2h15V7H12zm-8 6v6h6v-6H4zm2 2h2v2H6v-2zm6 0v2h15v-2H12zm-8 6v6h6v-6H4zm2 2h2v2H6v-2zm6 0v2h15v-2H12z"/>`),
		g.Group(children),
	)
}