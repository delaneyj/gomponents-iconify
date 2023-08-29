package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Duo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M24 5.5A18.5 18.5 0 1 0 42.5 24V8c0-1.8-.8-2.5-2.5-2.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="bevel" d="M16.5 17.1h11.3a2.65 2.65 0 0 1 2.6 2.6v8.7a2.65 2.65 0 0 1-2.6 2.6H16.5a2.65 2.65 0 0 1-2.6-2.6v-8.7a2.65 2.65 0 0 1 2.6-2.6Z"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="m30.4 24.9l4.6 4.7V18.4L30.4 23Z"/>`),
		g.Group(children),
	)
}