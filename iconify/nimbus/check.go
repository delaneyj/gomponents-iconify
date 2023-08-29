package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Check(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m15.12 2.23l-9.79 9.78L.88 7.56L0 8.44l5.33 5.34L16 3.11l-.88-.88z"/>`),
		g.Group(children),
	)
}