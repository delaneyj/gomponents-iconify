package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExponentialIdle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-5.141 5.948l-7.379 9.777m7.379 0l-7.379-9.777"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.143 33.444a6.017 6.017 0 0 1-5.285 3.108h0a6.235 6.235 0 0 1-6.217-6.217v-4.04a6.235 6.235 0 0 1 6.217-6.218h0a6.235 6.235 0 0 1 6.217 6.217v2.176H12.641"/>`),
		g.Group(children),
	)
}