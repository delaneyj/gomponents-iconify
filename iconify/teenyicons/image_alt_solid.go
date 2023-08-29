package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M11 4h-1V3h1v1Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v12.01A1.5 1.5 0 0 1 13.5 15h-12A1.5 1.5 0 0 1 0 13.5v-12Zm14 6.787l-2.5-2.498l-2.959 2.956L4.5 3.696L1 8.074V1.5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 .5.5v6.787Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}