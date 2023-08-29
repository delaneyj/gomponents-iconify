package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThinBigLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 5.5L0 12l6.5 6.5l.707-.707L1.914 12.5H24v-1H1.914l5.293-5.293L6.5 5.5Z"/>`),
		g.Group(children),
	)
}