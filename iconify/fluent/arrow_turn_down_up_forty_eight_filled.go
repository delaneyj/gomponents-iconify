package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnDownUpFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M38.866 6.88a1.5 1.5 0 0 0-1.943-.764l-12 5a1.5 1.5 0 0 0 1.154 2.769l8.781-3.659l-11.822 28.137L10.893 7.944a1.5 1.5 0 1 0-2.786 1.112L20.702 40.61c.83 2.078 3.76 2.104 4.627.042l12.293-29.258l3.513 7.728a1.5 1.5 0 0 0 2.73-1.241l-5-11Z"/>`),
		g.Group(children),
	)
}