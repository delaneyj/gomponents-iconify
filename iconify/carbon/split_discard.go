package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitDiscard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M28 16A12.01 12.01 0 0 0 17 4.05V2h-2v2.05a11.99 11.99 0 0 0 0 23.9V30h2v-2.05A12.01 12.01 0 0 0 28 16zM16 26V6a10 10 0 0 1 0 20z" fill="currentColor"/>`),
		g.Group(children),
	)
}