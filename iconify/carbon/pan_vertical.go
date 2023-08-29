package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m10 8l1.414 1.414L15 5.828v20.344l-3.586-3.586L10 24l6 6l6-6l-1.414-1.414L17 26.172V5.828l3.586 3.586L22 8l-6-6l-6 6z"/>`),
		g.Group(children),
	)
}