package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSmallDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m10.7 8.64l-2.5 2.5h-.7L5 8.64l.7-.71l1.65 1.64V4h1v5.57L10 7.92l.7.72z"/>`),
		g.Group(children),
	)
}