package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NaranjaX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.014 26.436l4.486 6.77M26.302 14.794L32.401 24m0 0l-6.099 9.206M38.5 14.794l-4.486 6.77M9.5 33.206V14.794l12.198 18.412V14.794"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1"/>`),
		g.Group(children),
	)
}