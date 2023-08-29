package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedDoubleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M14.5 7v8M.5 7v8M0 10.5h15m-15-3h15m-13-2h4m3 0h4M.5 6V.5h14V6"/>`),
		g.Group(children),
	)
}