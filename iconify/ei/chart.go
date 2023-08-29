package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M18 36h-2V26h-4v10h-2V24h8zm10 0h-2V20h-4v16h-2V18h8zm10 0h-2V14h-4v22h-2V12h8zM8 36h32v2H8z"/>`),
		g.Group(children),
	)
}