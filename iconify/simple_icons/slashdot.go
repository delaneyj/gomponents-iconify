package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slashdot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.777 0L7.037 24H1.868L14.605 0h5.172zm2.354 19.801a4.107 4.107 0 0 1-4.109 4.105a4.106 4.106 0 1 1 0-8.212a4.109 4.109 0 0 1 4.109 4.107z"/>`),
		g.Group(children),
	)
}