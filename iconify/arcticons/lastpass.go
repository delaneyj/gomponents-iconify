package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lastpass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.41 9.94A1.09 1.09 0 0 1 43.5 11v26a1.1 1.1 0 1 1-2.19 0V11a1.1 1.1 0 0 1 1.1-1.06ZM9.25 19.22A4.75 4.75 0 1 1 4.5 24a4.75 4.75 0 0 1 4.75-4.78Zm12.08 0A4.75 4.75 0 1 1 16.58 24a4.75 4.75 0 0 1 4.75-4.78Zm12.08 0A4.75 4.75 0 1 1 28.65 24a4.75 4.75 0 0 1 4.76-4.78Z"/>`),
		g.Group(children),
	)
}