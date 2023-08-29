package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tellurium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-26.772 9h12.588m-6.294 19v-19"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.65 31.103a4.748 4.748 0 0 1-4.128 2.397h0a4.75 4.75 0 0 1-4.75-4.75v-3.087a4.75 4.75 0 0 1 4.75-4.75h0a4.75 4.75 0 0 1 4.75 4.75v1.543h-9.5"/>`),
		g.Group(children),
	)
}