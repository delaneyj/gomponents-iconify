package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpLabelImportant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18.99h12.04L21 12l-4.97-7H4l5 7l-5 6.99z"/>`),
		g.Group(children),
	)
}