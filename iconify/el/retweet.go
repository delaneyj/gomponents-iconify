package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Retweet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M300 225L0 525h225v375h450L525 750H375V525h225L300 225zm225 75l150 150h150v225H600l300 300l300-300H975V300H525z"/>`),
		g.Group(children),
	)
}