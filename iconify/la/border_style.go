package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderStyle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h4v-2H7V7h18v2h2V5H5zm20 6v4h2v-4h-2zm0 6v4h2v-4h-2zm0 6v2h-2v2h4v-4h-2zm-14 2v2h4v-2h-4zm6 0v2h4v-2h-4z"/>`),
		g.Group(children),
	)
}