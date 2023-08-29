package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oversec(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="15.25" cy="17.78" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="32.75" cy="17.78" r="3.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.84 29l6.24-3.6l2.08 3.6L28 31.36l4.16 2.41l-2.08 3.6l-6.24-3.6l-6.24 3.6l-2.08-3.6l4.16-2.41L15.52 29l2.08-3.6Z"/>`),
		g.Group(children),
	)
}