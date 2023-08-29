package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FishLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M2 15c1.833-2.667 6.8-8 12-8c.923 0 1.754.105 2.5.287M2 9c1.833 2.667 6.8 8 12 8c.923 0 1.754-.105 2.5-.287m0 0C19.96 15.87 22 12 22 12s-2.04-3.87-5.5-4.713m0 9.426c-1-1.546-2.4-5.597 0-9.426M12 10.5c-.5.5-1.2 1.8 0 3m6-2.5h.001"/>`),
		g.Group(children),
	)
}