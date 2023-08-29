package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopySimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 66H40a6 6 0 0 0-6 6v144a6 6 0 0 0 6 6h144a6 6 0 0 0 6-6V72a6 6 0 0 0-6-6Zm-6 144H46V78h132Zm44-170v144a6 6 0 0 1-12 0V46H72a6 6 0 0 1 0-12h144a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}