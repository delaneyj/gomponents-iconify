package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Postid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.996 17.045v14.097m6.669-.188V16.858h2.29a7.069 7.069 0 0 1 7.049 7.048h0a7.07 7.07 0 0 1-7.048 7.048Z"/>`),
		g.Group(children),
	)
}