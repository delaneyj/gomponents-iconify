package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TallyMarkThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 19H7V5h2v14m4-14h-2v14h2V5m4 0h-2v14h2V5Z"/>`),
		g.Group(children),
	)
}