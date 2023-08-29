package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drumon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.8 23.1l38.7-12.4v-.5l-17.9.7c-3.8-2.2-5.1-.6-6 1.8l-15 9.8ZM23.15 40V21.18"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.68 17.67l.21 2.27l1.26 1.24l1.42-1.52l.2-2.96m-4.86-8.11l2.4-.92m-.19 3.45l-1.01-2.99"/>`),
		g.Group(children),
	)
}