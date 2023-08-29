package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sigma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m14.713 11.48l.694-1.48h.594l-1 6h-15v-1.16l5.18-6.113l-5.18-5.18V0h15.313l.688 4h-.537l-.293-.607C14.62 2.247 14.205 2 13.002 2H2.658l5.517 5.516l-4.647 5.483h8.474c1.813 0 2.291-.65 2.713-1.52z"/>`),
		g.Group(children),
	)
}