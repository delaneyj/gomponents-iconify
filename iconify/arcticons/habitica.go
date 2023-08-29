package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Habitica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.9 41.89c2.82 1 6-2.45 6-2.45a17.52 17.52 0 0 1-.36-4.29C19.49 29 26 26.88 26 26.88c-9.34 0-11.62-7.61-11.62-7.61s.47.37 2.78.37s2.74-1.38 2.74-1.38C8 19.06 8.68 5.5 8.68 5.5a43.75 43.75 0 0 0 6.81 6.5c5.29 4.08 8.32 2.27 11.94 7.93c0 0 .5-3.57 0-4.12s-3.73-1.32-3.9-2.81H26c1 0 3.74-2.69 5.77-2.69s1.54.71 2.47.71S38 9.49 38 14.82c0 0-3-.33-3 1.38s4.8 1.2 4.8 7.8h-1.54a15.34 15.34 0 0 1 1.21 4.72c0 1.93-2.26 4.34-2.26 4.34v5.88c0 1.21.44 1.65 1.21 1.65a9.16 9.16 0 0 0 2.14-.27c.87-.24 2.71 2.18.55 2.18h-6.88c-.48 0-.81-.15-.81-.59s2-1.23 2-2.24s-1.23-4.1-3.42-4.1s-4.86 2.09-4.86 4.07c0 1.08.66 1.1 1.34 1.1s1.34-.48 2-.48s1.41 1.26 1.41 2.24H19c-1.56 0-8.13.33-8.13-4.62s5.29-4 5.29-7.44a1.56 1.56 0 0 0-1.45-1.75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.77 24.44v-3.67h-3.1v6.1h1.6v1.82h4.45v-4.25h-2.95zM5.88 17.9h2.79v2.88H5.88z"/>`),
		g.Group(children),
	)
}