package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThinLeftBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m1.124 15.922l-.017-8.721c0-.648.611-1.029.895-1.027c.57.001.948.464.949 1.034l.01 5.347L15.355.355a.938.938 0 0 1 1.327.018A.936.936 0 0 1 16.7 1.7L4.554 14.012l5.333.008c.57.002 1.033.465 1.033 1.033c.003.57-.457.887-1.027.885l-8.769-.016z"/>`),
		g.Group(children),
	)
}