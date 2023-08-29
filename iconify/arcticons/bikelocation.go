package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bikelocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5h0A15.93 15.93 0 0 0 8.08 20.42c0 6.22 3.93 11.9 7.8 16a51.8 51.8 0 0 0 7.73 6.78l.39.27l.39-.27a51.8 51.8 0 0 0 7.73-6.78c3.87-4.14 7.8-9.81 7.8-16A15.93 15.93 0 0 0 24 4.5Z"/><circle cx="31.41" cy="20" r="3.98" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="16.64" cy="20" r="3.98" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.12 13.06L23.72 20h-7.08l5.42-6.94h7.06"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.41 20l-3.14-9.52h-2.31m-4.42.38L23.72 20m-3.59-9.14h2.94M9.86 27.77h28.28"/>`),
		g.Group(children),
	)
}