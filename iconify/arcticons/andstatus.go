package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Andstatus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.93 40.88C28 37.67 15 22.2 15 22.2c-1.64-2.15-4.53-5-4.53-9.18a7.5 7.5 0 0 1 7.7-7.42a7.36 7.36 0 0 1 7.41 7.4c0 4.2-3.4 7.78-9.69 9C8.88 23.33 4.5 27.17 4.5 33.51c0 5.13 3.23 8.89 9.69 8.89c8.49 0 13.23-8.83 18.11-16"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.05 7.4l-5.73 4.79l5.75 7.64M36 24.45a4.48 4.48 0 0 0-4.47 4.49h0A4.48 4.48 0 0 0 36 33.42h1.5m0 0H39a4.48 4.48 0 0 1 4.48 4.49h0A4.48 4.48 0 0 1 39 42.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.07 26c-1.23-1-2.57-1.51-5.57-1.51H36m-4.07 16.39c1.24 1 2.58 1.52 5.57 1.52H39"/>`),
		g.Group(children),
	)
}