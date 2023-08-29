package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v3a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V8Zm7 1.5a.5.5 0 0 0 .5.5H14a.5.5 0 0 0 0-1H9.5a.5.5 0 0 0-.5.5Zm-1 0a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Z"/>`),
		g.Group(children),
	)
}