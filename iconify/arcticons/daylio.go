package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Daylio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.11 32.44l21.78.2c-1.07 10.49-22.12 8.64-21.78-.2Z"/><ellipse cx="12.11" cy="22.98" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.75" ry="2.39"/><ellipse cx="35.89" cy="22.98" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.75" ry="2.39"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}