package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopySimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 68H40a4 4 0 0 0-4 4v144a4 4 0 0 0 4 4h144a4 4 0 0 0 4-4V72a4 4 0 0 0-4-4Zm-4 144H44V76h136Zm40-172v144a4 4 0 0 1-8 0V44H72a4 4 0 0 1 0-8h144a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}