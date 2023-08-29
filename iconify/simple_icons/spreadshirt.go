package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spreadshirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6.306L7.796 2.102L0 9.898l12 12l12-12l-7.796-7.796zm0 12L3.592 9.898l4.204-4.204L12 9.898l4.184-4.184l4.204 4.204"/>`),
		g.Group(children),
	)
}