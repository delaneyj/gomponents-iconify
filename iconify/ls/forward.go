package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M494 38v124C-36 162 1 656 1 656c104-345 493-245 493-245v127l274-254z"/>`),
		g.Group(children),
	)
}