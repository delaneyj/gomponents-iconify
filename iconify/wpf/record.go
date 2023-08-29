package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Record(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M13 6.188a6.812 6.812 0 1 0 0 13.625a6.812 6.812 0 1 0 0-13.625z"/>`),
		g.Group(children),
	)
}