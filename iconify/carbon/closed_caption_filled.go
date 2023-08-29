package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptionFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 6H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2Zm-13 7h-5v6h5v2h-5a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h5Zm10 0h-5v6h5v2h-5a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h5Z"/>`),
		g.Group(children),
	)
}