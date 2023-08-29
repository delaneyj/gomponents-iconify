package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 4.25A.25.25 0 0 1 2.25 4h10.5a.25.25 0 0 1 .25.25v7.25H2V4.25ZM2.25 3C1.56 3 1 3.56 1 4.25V12H0v.5a.5.5 0 0 0 .5.5h14a.5.5 0 0 0 .5-.5V12h-1V4.25C14 3.56 13.44 3 12.75 3H2.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}