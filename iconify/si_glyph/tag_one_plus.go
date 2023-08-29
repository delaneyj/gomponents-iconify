package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagOnePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m13.613 1.047l-4.502.003l-7.779 7.77a1.14 1.14 0 0 0 0 1.607l5.244 5.242a1.135 1.135 0 0 0 1.605 0l7.787-7.707l.002-4.534l-2.357-2.381zM4.984 9.014l1.023-1.023l-.941-.935l.973-.968l.939.938l1.025-1.021l.988.985l-1.024 1.021l.94.937l-.973.967l-.938-.937l-1.024 1.02l-.988-.984zm2.137 4.543l-.789-.788l4.243-4.243l-.727-.725l.786-.788l1.514 1.514l-5.027 5.03zM15.942.36h.984v1.796h-.984z"/>`),
		g.Group(children),
	)
}