package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pokitmeter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.7 5.7v5.9a12.86 12.86 0 0 0-3.8 23.1l3.8-3.9v11.5C.2 38.3.6 9.7 20.7 5.7Zm6.6 36.6v-5.9a12.86 12.86 0 0 0 3.8-23.1l-3.8 3.9V5.7c20.5 4 20.1 32.6 0 36.6Z"/>`),
		g.Group(children),
	)
}