package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTextJustifySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 4H0V3h15v1Zm0 4H0V7h15v1Zm0 4H0v-1h15v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}