package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Four(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M398 0v494h89v72h-89v209h-71V566H0zM138 494h189V227z"/>`),
		g.Group(children),
	)
}