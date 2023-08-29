package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m9.984 6.943l-2.926-.671V5.014h2.926v1.929zm0 2.997l-2.926-.67V8.011l2.926.671V9.94zm0 3.029l-2.926-.671v-1.259l2.926.672v1.258zm-1.18 3.029h-.566l-1.18-1.93l2.926.671l-1.18 1.259zM9 .099v1.933H7.938V.095c-1.656.21-2.926 1.337-2.926 2.706c0 1.521 6.957 1.521 6.957 0C11.969 1.438 10.647.315 9 .099z"/>`),
		g.Group(children),
	)
}