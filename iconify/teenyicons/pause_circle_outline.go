package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M6.5 5v5m2-5v5m-1 4.5a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		g.Group(children),
	)
}