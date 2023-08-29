package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongBottomDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.021 15.55L7.014 12H5v7h7v-2.014l-3.55-.007L19 6.429L17.571 5L7.021 15.55Z"/>`),
		g.Group(children),
	)
}