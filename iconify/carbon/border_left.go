package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 29V3h2v26zm4 0v-2h2v2zm4 0v-2h2v2zm4 0v-2h2v2zm4 0v-2h2v2zm4 0v-2h2v2zm4 0v-2h2v2zm0-4v-2h2v2zm0-4v-2h2v2zm0-4v-2h2v2zm0-8V7h2v2zm0 4v-2h2v2zM7 5V3h2v2zm4 0V3h2v2zm4 0V3h2v2zm4 0V3h2v2zm4 0V3h2v2zm4 0V3h2v2zM8 10h10v2H8zm0 5h6v2H8z"/>`),
		g.Group(children),
	)
}