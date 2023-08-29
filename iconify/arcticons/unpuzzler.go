package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unpuzzler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.63 17.86a5.82 5.82 0 0 0-4.57 2.22V9.72H22.43a5.87 5.87 0 1 1-7.3 0H4.5v28.56h10.63a5.87 5.87 0 1 1 7.3 0h10.63v-10.9a5.86 5.86 0 1 0 4.57-9.52Z"/>`),
		g.Group(children),
	)
}