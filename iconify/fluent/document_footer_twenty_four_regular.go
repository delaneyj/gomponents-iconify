package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentFooterTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.501 16a1.5 1.5 0 1 0 0 3h7a1.5 1.5 0 0 0 0-3h-7Zm11.49-11.908a2.25 2.25 0 0 0-2.245-2.096h-11.5l-.154.005a2.25 2.25 0 0 0-2.096 2.245v15.498l.005.154a2.25 2.25 0 0 0 2.245 2.096h11.5l.154-.005a2.25 2.25 0 0 0 2.096-2.245V4.246l-.005-.154ZM6.246 3.496h11.5l.102.007a.75.75 0 0 1 .648.743v15.498l-.007.102a.75.75 0 0 1-.743.648h-11.5l-.102-.007a.75.75 0 0 1-.648-.743V4.246l.007-.102a.75.75 0 0 1 .743-.648Z"/>`),
		g.Group(children),
	)
}