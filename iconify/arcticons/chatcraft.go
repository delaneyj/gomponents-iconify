package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chatcraft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.09 17.03L23.91 10L11.8 16.99l12.18 7.03l12.11-6.99z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.98 24.02L11.8 16.99v13.98L23.98 38V24.02zm12.11-6.99l-12.11 6.99V38l12.11-6.99V17.03z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.98 30.94l-2.5-4.68l-7.18-4.15l-2.5 1.8m24.29.04l-2.48-1.79l-7.14 4.13l-2.49 4.65M31.35 17l-7.43-4.29l-7.38 4.27l7.43 4.28L31.35 17zm-3.71-2.14l-7.38 4.26m7.36.03l-7.42-4.28"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.52 21.52 0 0 0 5.15 34.36L2.5 45.5l11.14-2.65A21.5 21.5 0 1 0 24 2.5Z"/>`),
		g.Group(children),
	)
}