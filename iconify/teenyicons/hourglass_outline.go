package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.5 0v4.672a2 2 0 0 0 .586 1.414l.707.707a1 1 0 0 1 0 1.414l-1 1a1 1 0 0 0-.293.707V15m8-15v5.086a1 1 0 0 1-.293.707l-1 1a1 1 0 0 0 0 1.414l1 1a1 1 0 0 1 .293.707V15M1 .5h13m-13 14h13"/>`),
		g.Group(children),
	)
}