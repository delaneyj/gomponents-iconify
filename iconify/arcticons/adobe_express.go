package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeExpress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4" ry="4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.21 34a2.97 2.97 0 0 1-2.713-1.76L24 22.15l-2.635 5.912h.838a2.969 2.969 0 0 1 0 5.937H16.79a2.97 2.97 0 0 1-2.711-4.178l5.886-13.202A4.423 4.423 0 0 1 24 13.999h0a4.422 4.422 0 0 1 4.035 2.62l5.886 13.203A2.969 2.969 0 0 1 31.21 34Z"/>`),
		g.Group(children),
	)
}