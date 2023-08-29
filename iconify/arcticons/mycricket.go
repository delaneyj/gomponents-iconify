package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mycricket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.86 33.81l7.84 8.4h11.66l-13.61-14.3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.71 43l-.06-11.67l17-16.71H38Zm-.07-11.67L9.82 5h8.12v18.16"/>`),
		g.Group(children),
	)
}