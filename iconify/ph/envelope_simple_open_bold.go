package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeSimpleOpenBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m230.66 86l-96-64a12 12 0 0 0-13.32 0l-96 64A12 12 0 0 0 20 96v104a20 20 0 0 0 20 20h176a20 20 0 0 0 20-20V96a12 12 0 0 0-5.34-10ZM128 46.42l74.86 49.91L141.61 140h-27.22L53.14 96.33ZM44 196v-76.71l59.58 42.48a12 12 0 0 0 7 2.23h34.9a12 12 0 0 0 7-2.23L212 119.29V196Z"/>`),
		g.Group(children),
	)
}