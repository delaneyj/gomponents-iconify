package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Collaboraoffice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.46 34.61L18 43.06a1.48 1.48 0 0 0 2.11 0h0l18-18a1.49 1.49 0 0 0 0-2.11h0l-18-18a1.49 1.49 0 0 0-2.12 0h0l-8.5 8.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.46 34.61l9.36-9.36a1.73 1.73 0 0 0 0-2.45h0l-9.36-9.36Z"/>`),
		g.Group(children),
	)
}