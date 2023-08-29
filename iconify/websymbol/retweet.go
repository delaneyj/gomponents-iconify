package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Retweet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1120 481h240l-300 320l-300-320h240V241H600L483 121h636zM760 681l117 120H241l-1-360H0l300-320l300 320H360v240h400z"/>`),
		g.Group(children),
	)
}