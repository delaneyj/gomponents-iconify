package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudFoundryTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 11V9h-8v14h2v-6h5v-2h-5v-4h6zM15 23H9a2 2 0 0 1-2-2V11a2 2 0 0 1 2-2h6v2H9v10h6z"/>`),
		g.Group(children),
	)
}