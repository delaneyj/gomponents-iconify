package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Techcombank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 24.02L30.514 11.355L24 17.688l6.516 6.333L24 30.353l6.516 6.293L43.5 24.02ZM24 17.689l-6.472 6.333L24 30.353l-6.47 6.292L4.5 24.021l13.028-12.667L24 17.688"/>`),
		g.Group(children),
	)
}