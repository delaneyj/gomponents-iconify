package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RibbonTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.002 15.244L17 21.245a.75.75 0 0 1-1.182.613L12 19.171l-3.817 2.687a.75.75 0 0 1-1.181-.613l-.002-6A7.966 7.966 0 0 0 12 17a7.966 7.966 0 0 0 5.002-1.756ZM12 2a7 7 0 1 1 0 14a7 7 0 0 1 0-14Z"/>`),
		g.Group(children),
	)
}