package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M25.25 5.25a1.25 1.25 0 1 0-2.5 0v32.446L10.141 24.874a1.25 1.25 0 1 0-1.782 1.752l14.75 15a1.25 1.25 0 0 0 1.782 0l14.75-15a1.25 1.25 0 1 0-1.782-1.752L25.25 37.696V5.25Z"/>`),
		g.Group(children),
	)
}