package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Daserste(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M39.57 5.5v29.73L19.78 42.5V21.32L8.43 25.5V17Z"/>`),
		g.Group(children),
	)
}