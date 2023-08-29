package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloudera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12S18.627 0 12 0zm.344 20.251a8.25 8.25 0 1 1 0-16.502a8.21 8.21 0 0 1 5.633 2.234L15.519 8.53a4.686 4.686 0 0 0-3.175-1.239a4.709 4.709 0 1 0 3.284 8.081l2.657 2.346a8.224 8.224 0 0 1-5.941 2.533z"/>`),
		g.Group(children),
	)
}