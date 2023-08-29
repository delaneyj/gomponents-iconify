package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.45 7.021L12 7.014V5H5v7h2.014l.007-3.55L17.571 19L19 17.571L8.45 7.021Z"/>`),
		g.Group(children),
	)
}