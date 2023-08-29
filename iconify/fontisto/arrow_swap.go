package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSwap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m0 16.16l5.746-5.746v4.214H24v3.065H5.746v4.214zm18.254-7.28H0V5.814h18.254V1.6L24 7.346l-5.746 5.746z"/>`),
		g.Group(children),
	)
}