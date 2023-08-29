package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataAreaTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 3.75a.75.75 0 0 1 1.5 0V19.5h15.75a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1-.75-.75V3.75Zm16.5 3a.75.75 0 0 0-1.2-.6l-5.6 4.2l-3.82-2.246a.75.75 0 0 0-.715-.025L5.5 9.411V18.5h14V6.75Z"/>`),
		g.Group(children),
	)
}