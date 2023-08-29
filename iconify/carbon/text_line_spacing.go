package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextLineSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 6h13v2H17zm0 6h10v2H17zm0 6h13v2H17zm0 6h10v2H17zm-5.41-10.59L8 9.83l-3.59 3.59L3 12l5-5l5 5l-1.41 1.41zm0 5.18L8 22.17l-3.59-3.59L3 20l5 5l5-5l-1.41-1.41z"/>`),
		g.Group(children),
	)
}