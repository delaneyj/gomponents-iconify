package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareRightLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 34H48a14 14 0 0 0-14 14v160a14 14 0 0 0 14 14h160a14 14 0 0 0 14-14V48a14 14 0 0 0-14-14Zm2 174a2 2 0 0 1-2 2H48a2 2 0 0 1-2-2V48a2 2 0 0 1 2-2h160a2 2 0 0 1 2 2Zm-37.76-84.24a6 6 0 0 1 0 8.48l-32 32a6 6 0 0 1-8.48-8.48L153.51 134H88a6 6 0 0 1 0-12h65.51l-21.75-21.76a6 6 0 0 1 8.48-8.48Z"/>`),
		g.Group(children),
	)
}