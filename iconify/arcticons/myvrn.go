package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myvrn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.903 24l4.798-8.636h9.595L15.498 24l4.798 8.636h-9.595Zm21.801 18.5h9.595l4.798-8.636l-4.798-8.636l-4.797 8.636h-9.596Zm4.798-28.364l4.797 8.636l4.798-8.636L37.299 5.5h-9.595l-4.798 8.636Z"/>`),
		g.Group(children),
	)
}