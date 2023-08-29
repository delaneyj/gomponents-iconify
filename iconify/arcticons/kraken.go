package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kraken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.877 28.116a2.627 2.627 0 0 1 5.251 0h0m-15.754 0a2.627 2.627 0 0 1 5.252 0m-15.754 0a2.582 2.582 0 0 1 2.625-2.535h0a2.582 2.582 0 0 1 2.626 2.535m5.251 8.009a2.52 2.52 0 0 1-1.313 2.195a2.708 2.708 0 0 1-2.625 0a2.52 2.52 0 0 1-1.313-2.196m26.257.001a2.52 2.52 0 0 1-1.313 2.195a2.707 2.707 0 0 1-2.626 0a2.52 2.52 0 0 1-1.313-2.196m-5.251.001a2.52 2.52 0 0 1-1.313 2.195a2.708 2.708 0 0 1-2.626 0a2.52 2.52 0 0 1-1.312-2.196"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.62 25.818c0-9.1 8.23-16.478 18.38-16.478s18.38 7.378 18.38 16.478M10.872 36.125a2.52 2.52 0 0 1-1.313 2.195a2.707 2.707 0 0 1-2.626 0a2.52 2.52 0 0 1-1.313-2.195m0-10.307v10.307m36.76-10.307v10.307m-5.252 0v-8.01m-5.251.001v8.009m-5.251 0v-8.01m-5.252.001v8.009m-5.251 0v-8.01m-5.251.001v8.009"/>`),
		g.Group(children),
	)
}