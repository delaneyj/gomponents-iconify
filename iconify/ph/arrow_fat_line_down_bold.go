package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLineDownBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M235.09 131.41A12 12 0 0 0 224 124h-36V80a12 12 0 0 0-12-12H80a12 12 0 0 0-12 12v44H32a12 12 0 0 0-8.48 20.49l96 96a12 12 0 0 0 17 0l96-96a12 12 0 0 0 2.57-13.08ZM128 215l-67-67h19a12 12 0 0 0 12-12V92h72v44a12 12 0 0 0 12 12h19ZM68 40a12 12 0 0 1 12-12h96a12 12 0 0 1 0 24H80a12 12 0 0 1-12-12Z"/>`),
		g.Group(children),
	)
}