package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VaultLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 42H40a14 14 0 0 0-14 14v136a14 14 0 0 0 14 14h18v18a6 6 0 0 0 12 0v-18h116v18a6 6 0 0 0 12 0v-18h18a14 14 0 0 0 14-14V56a14 14 0 0 0-14-14Zm0 152H40a2 2 0 0 1-2-2V56a2 2 0 0 1 2-2h176a2 2 0 0 1 2 2v66h-20.4a46 46 0 1 0 0 12H218v58a2 2 0 0 1-2 2Zm-51.37-72a14 14 0 1 0 0 12h20.83a34 34 0 1 1 0-12Z"/>`),
		g.Group(children),
	)
}