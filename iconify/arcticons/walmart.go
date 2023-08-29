package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Walmart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 15.42V4.5m7.41 15.2l9.46-5.46m-9.46 14.02l9.46 5.46M24 32.58V43.5m-7.41-15.24l-9.46 5.46m9.46-14.02l-9.46-5.46"/>`),
		g.Group(children),
	)
}