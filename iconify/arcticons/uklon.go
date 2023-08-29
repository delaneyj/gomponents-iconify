package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uklon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.696 14.778c0-5.736 32.5-12.427 35.369-7.647c3.823 5.735-1.912 35.369-7.648 35.369c-6.691 0-27.722-20.074-27.722-27.722h0Z"/>`),
		g.Group(children),
	)
}