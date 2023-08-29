package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4.688L3.781 16.905l1.438 1.407L16 7.53l10.781 10.782l1.438-1.407zm0 7L3.781 23.905l1.438 1.407L16 14.53l10.781 10.781l1.438-1.406z"/>`),
		g.Group(children),
	)
}