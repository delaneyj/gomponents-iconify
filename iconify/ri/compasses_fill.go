package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassesFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 4.123V1.997h2v2.126a4.002 4.002 0 0 1 1.802 6.728l6.064 10.502l-1.732 1l-6.063-10.501a4.004 4.004 0 0 1-2.142 0L4.866 22.353l-1.732-1l6.063-10.502A4.002 4.002 0 0 1 11 4.123Zm1 4.874a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}