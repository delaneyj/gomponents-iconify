package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 1A1.5 1.5 0 0 0 1 2.5V5H0V2.5A2.5 2.5 0 0 1 2.5 0H5v1H2.5Zm10 0H10V0h2.5A2.5 2.5 0 0 1 15 2.5V5h-1V2.5A1.5 1.5 0 0 0 12.5 1Zm.5 7H2V7h11v1ZM0 12.5V10h1v2.5A1.5 1.5 0 0 0 2.5 14H5v1H2.5A2.5 2.5 0 0 1 0 12.5Zm14 0V10h1v2.5a2.5 2.5 0 0 1-2.5 2.5H10v-1h2.5a1.5 1.5 0 0 0 1.5-1.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}