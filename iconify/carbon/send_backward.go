package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 7H2V4a2.002 2.002 0 0 1 2-2h3v2H4zm3 15H4a2.002 2.002 0 0 1-2-2v-3h2v3h3zM2 10h2v4H2zm20-3h-2V4h-3V2h3a2.002 2.002 0 0 1 2 2zM10 2h4v2h-4zm18 28H12a2.002 2.002 0 0 1-2-2V12a2.002 2.002 0 0 1 2-2h16a2.002 2.002 0 0 1 2 2v16a2.002 2.002 0 0 1-2 2zM12 12v16h16V12z"/>`),
		g.Group(children),
	)
}