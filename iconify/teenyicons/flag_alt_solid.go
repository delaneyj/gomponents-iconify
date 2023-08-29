package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2.254.065a.5.5 0 0 1 .503.006l10 6a.5.5 0 0 1-.033.876L3 11.81V15H2V.5a.5.5 0 0 1 .254-.435Z"/>`),
		g.Group(children),
	)
}