package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XosLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.2 28.435c.696.907 1.57 1.245 2.785 1.245h1.681a2.834 2.834 0 0 0 2.834-2.834v-.012A2.834 2.834 0 0 0 35.666 24h-1.855a2.837 2.837 0 0 1-2.837-2.837h0a2.843 2.843 0 0 1 2.843-2.843h1.673c1.215 0 2.088.338 2.785 1.245"/><rect width="7.526" height="11.359" x="21.16" y="18.32" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.763" ry="3.763"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.427 18.32L13.963 24l4.464 5.68M9.5 18.32l2.764 3.518m0 4.324L9.5 29.68"/>`),
		g.Group(children),
	)
}