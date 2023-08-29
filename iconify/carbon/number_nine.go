package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 9h-4a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h4v3h-5v2h5a2 2 0 0 0 2-2V11a2 2 0 0 0-2-2Zm0 7h-4v-5h4Z"/>`),
		g.Group(children),
	)
}