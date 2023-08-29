package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDobuleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.4 12L22.378 0L24.8 2.43L15.253 12l9.547 9.57L22.378 24zM0 12L11.978 0L14.4 2.43L4.853 12l9.547 9.57L11.978 24z"/>`),
		g.Group(children),
	)
}