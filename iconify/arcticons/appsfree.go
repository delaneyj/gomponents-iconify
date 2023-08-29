package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Appsfree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.539 28.849h12.34m-25.072 0h10.752m1.98 9.52V9.631H42.5m-37 28.738l17.059-25.75"/>`),
		g.Group(children),
	)
}