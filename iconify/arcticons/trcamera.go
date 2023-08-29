package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trcamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10 22.5l-2.55 1.78s-2-2.59-1.67-3.56c1.28-3.78 9.65-5.87 9.65-5.87a9.83 9.83 0 0 0-3.72-.43c1.78-2.05 8.09-4.8 8.09-4.8c-.49-3.18 2.1-5.12 3.07-5.12s5.44 4.74 5.44 4.74a26.08 26.08 0 0 1 3.61 1.19c-2.31.05-4.42 2.64-4.42 2.64c2.48-1.13 8.31 0 8.52 0a16.87 16.87 0 0 1 1.89 2.21A7.6 7.6 0 0 0 33 16.41c2.7-.48 6.42 1.3 6.85 1.73c.63.63 7.21 13.05-4.53 21.39c-13.91 9.89-23.39-1.83-23.39-1.83s11.75 9.27 22.42.43c9.87-8.17 2.86-17.06 2.86-17.06c5 13.75-10.35 18.6-10.35 18.6c9.51-5.23 6.63-14 6.63-14c0 7.66-6.9 10.33-6.9 10.33a4.11 4.11 0 0 0 .47-1.77a3.29 3.29 0 0 0-2.3-3.23s-2.19 4.85-2.35 5s-.24.44-3.15.44A5.06 5.06 0 0 1 15 34c1.54-.28 4.41-2 4.41-4.29a3.66 3.66 0 0 0-3.68-3.59c-1.7 0-4.36 1.57-5.13 2L7.9 24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10 27.17c.09-2.94 1.79-6.18 6.48-8.24s11.28-.48 12.49.45c0 0-3.64.08-4.37 1.41c0 0 5.38 3 5.38 7.48s-2.91 6-2.91 6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13 28.69a2 2 0 0 1-.59-1.57m2.93 5.27a2.88 2.88 0 0 0 .85 1.18"/>`),
		g.Group(children),
	)
}