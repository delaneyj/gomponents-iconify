package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockwiseSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.295 2.002L6.147.854l.706-.708L9.207 2.5L6.853 4.85l-.706-.707l1.141-1.141A5.499 5.499 0 0 0 2 8.495a5.499 5.499 0 0 0 5.5 5.496A5.5 5.5 0 0 0 13 8.495v-.5h1v.5a6.499 6.499 0 0 1-6.5 6.496A6.499 6.499 0 0 1 1 8.495a6.499 6.499 0 0 1 6.295-6.493Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}