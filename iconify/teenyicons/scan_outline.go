package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 5V2.5a2 2 0 0 1 2-2H5m5 0h2.5a2 2 0 0 1 2 2V5m-14 5v2.5a2 2 0 0 0 2 2H5m9.5-4.5v2.5a2 2 0 0 1-2 2H10m-8-7h11"/>`),
		g.Group(children),
	)
}