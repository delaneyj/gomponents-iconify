package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.05 19.55a9.06 9.06 0 0 0-9.54-9c-4.89.26-8.56 4.65-8.56 9.55v8.39A9 9 0 0 0 24 37.5h0a9 9 0 0 0 9-9v-8.9m-2.1-5.9L17.09 34.29"/>`),
		g.Group(children),
	)
}