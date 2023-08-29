package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpPanTool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 4v20H10.02L1 14.83L2.9 13L8 15.91V3h3v8h1V0h3v11h1V1h3v10h1V4h3z"/>`),
		g.Group(children),
	)
}