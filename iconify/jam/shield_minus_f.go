package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldMinusF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.649 3.322L8 .565l7.351 2.757a1 1 0 0 1 .649.936v4.307c0 3.177-1.372 6.2-3.763 8.292L8 20.565l-4.237-3.708A11.019 11.019 0 0 1 0 8.565V4.258a1 1 0 0 1 .649-.936zM5 8.565a1 1 0 1 0 0 2h6a1 1 0 0 0 0-2H5z"/>`),
		g.Group(children),
	)
}