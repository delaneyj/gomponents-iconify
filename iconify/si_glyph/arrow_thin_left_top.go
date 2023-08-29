package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThinLeftTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.885.006c.573-.002 1.075.403 1.073.974c0 .568-.507.942-1.079.944l-5.354.008l12.193 12.312a.934.934 0 0 1-.018 1.327a.948.948 0 0 1-1.334.018l-12.44-12.2l-.011 5.569c-.001.57-.381 1.033-.952 1.034c-.285.002-.898-.379-.898-1.027L1.082.022L9.885.006z"/>`),
		g.Group(children),
	)
}