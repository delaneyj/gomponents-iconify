package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 4v16h6V4zm2 2h2v12h-2zm-2 16v6h6v-6zm2 2h2v2h-2z"/>`),
		g.Group(children),
	)
}