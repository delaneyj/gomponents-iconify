package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 4v16h6V4h-6zm2 2h2v12h-2V6zm-2 16v6h6v-6h-6zm2 2h2v2h-2v-2z"/>`),
		g.Group(children),
	)
}