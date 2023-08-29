package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Esea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m18.74 4l-6.808 8.995L0 13.178l7.776 5.959l-6.385 9.171l11.317-5.447l7.615 5.765v-9.489L32 13.37L8.995 18.735c.792-.797 1.5-1.672 2.12-2.609c.407-.74.688-1.541.817-2.38l7.459-.204z"/>`),
		g.Group(children),
	)
}