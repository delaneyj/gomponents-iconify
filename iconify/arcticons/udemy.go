package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Udemy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.858 15.824v18.335A9.305 9.305 0 0 0 24 43.5a9.305 9.305 0 0 0 9.142-9.34V15.823m-18.273-5.607L24 4.5l9.131 5.716"/>`),
		g.Group(children),
	)
}