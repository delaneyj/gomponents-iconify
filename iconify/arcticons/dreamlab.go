package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dreamlab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.789 35.514a8.288 8.288 0 1 1 5.328-14.638"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.116 20.876a9.946 9.946 0 1 1 19.491 3.893m.006-.007a5.387 5.387 0 1 1 .5 10.752m-25.324 0h25.324"/>`),
		g.Group(children),
	)
}