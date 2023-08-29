package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewGridOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 5.5h15m-15-4h15m-15 8h15m-15 4h15M9.5 0v15m4-15v15m-8-15v15m-4-15v15"/>`),
		g.Group(children),
	)
}