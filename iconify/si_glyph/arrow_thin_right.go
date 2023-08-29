package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThinRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m16.992 7.98l-6.305 6.693c-.459.459-1.16.296-1.359.094c-.402-.403-.342-.998.061-1.402l3.895-4.438L1.97 8.925a.94.94 0 0 1-.927-.952a.937.937 0 0 1 .927-.95l11.212-.013l-3.855-4.271a1.036 1.036 0 0 1 0-1.461c.4-.405.95-.304 1.352.101l6.313 6.601z"/>`),
		g.Group(children),
	)
}