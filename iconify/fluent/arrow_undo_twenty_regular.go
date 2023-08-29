package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUndoTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 2.5a.5.5 0 0 0-1 0v4.9a.6.6 0 0 0 .6.6h4.9a.5.5 0 0 0 0-1H5.906l3.473-3.019a4 4 0 0 1 5.248 6.038l-8.172 7.104a.5.5 0 0 0 .656.754l8.172-7.103a5 5 0 0 0-6.56-7.547L5 6.463V2.5Z"/>`),
		g.Group(children),
	)
}