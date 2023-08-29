package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 6h20v2H6zm4 6h12v2H10zm-4 6h20v2H6zm4 6h12v2H10z"/>`),
		g.Group(children),
	)
}