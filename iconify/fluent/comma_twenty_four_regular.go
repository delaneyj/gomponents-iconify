package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommaTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.62 12.023a4 4 0 1 1 1.354-3.484c.356 1.81.352 3.967-.464 5.9c-.862 2.04-2.601 3.76-5.57 4.537a.75.75 0 1 1-.38-1.451c2.531-.663 3.892-2.07 4.568-3.67a7.224 7.224 0 0 0 .492-1.832Z"/>`),
		g.Group(children),
	)
}