package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Haven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c9.043-3.117 15.489-10.363 16.5-19.589a79.36 79.36 0 0 0-.071-12.027a2.541 2.541 0 0 0-2.468-2.366c-4.091-.126-8.846-.808-12.52-4.427a2.052 2.052 0 0 0-2.881 0c-3.675 3.619-8.43 4.301-12.52 4.427a2.541 2.541 0 0 0-2.468 2.366A79.36 79.36 0 0 0 7.5 23.911C8.511 33.137 14.957 40.383 24 43.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.372 18.156L24 12.482l-9.372 5.674v13.86h18.744v-13.86zm-18.744 0l-2.534 1.603m21.278-1.603l2.535 1.603"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.302 26.017h3.395v6h-3.395z"/>`),
		g.Group(children),
	)
}