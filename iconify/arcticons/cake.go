package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.977 24c0 3.998 1.922 6.288-.078 9.75s-4.946 2.943-8.41 4.942S28 43.5 24 43.5s-5.024-2.81-8.488-4.808S9.1 37.212 7.1 33.75s-.078-5.752-.078-9.75s-1.922-6.288.078-9.75s4.946-2.943 8.41-4.942S20 4.5 24 4.5s5.024 2.81 8.488 4.808s6.411 1.48 8.411 4.942s.078 5.752.078 9.75Z"/>`),
		g.Group(children),
	)
}