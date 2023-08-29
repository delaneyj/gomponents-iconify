package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottomSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 0H3v12h4V0Zm5 4H8v8h4V4Zm3 10H0v1h15v-1Z"/>`),
		g.Group(children),
	)
}