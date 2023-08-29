package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Doodle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.8 4.5h8.1v39h-8.1v-1.8c-5.4 3.4-12.5 1.8-15.9-3.6s-1.8-12.5 3.6-15.9c3.7-2.4 8.5-2.4 12.2 0V4.5zm-6.1 22.7c-2.6 0-4.8 2.1-4.8 4.8s2.1 4.8 4.8 4.8s4.8-2.1 4.8-4.8c0-2.6-2.2-4.8-4.8-4.8c0 0 0 0 0 0z"/>`),
		g.Group(children),
	)
}