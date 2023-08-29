package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cutlery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m114 221l-89-90Q0 106 0 71t25-60l150 149zm145-39l-31 31l146 147l-30 30l-146-147L51 390l-31-30l209-208q-12-24-4-56t33-57q31-30 69-35t61.5 18.5T407 84t-36 69q-25 25-56.5 33t-55.5-4z"/>`),
		g.Group(children),
	)
}