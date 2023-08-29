package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DividerLineSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2H3V1h9v1Zm-2 3H3V4h7v1Zm5 3H0V7h15v1Zm-3 3H3v-1h9v1Zm-2 3H3v-1h7v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}