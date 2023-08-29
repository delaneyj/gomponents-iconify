package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vibrons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 17.044v14.23m7.4 10.206V5.969m7.4 28.391V6.106m7.4 7.13v28.796m7.4-.397V6.292m7.4 5.564v21.29"/>`),
		g.Group(children),
	)
}